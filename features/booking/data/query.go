package data

import (
	"KosKita/features/booking"
	kd "KosKita/features/kos/data"
	"KosKita/utils/externalapi"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type bookingQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func NewBooking(db *gorm.DB, mid externalapi.MidtransInterface) booking.BookDataInterface {
	return &bookingQuery{
		db:              db,
		paymentMidtrans: mid,
	}
}

// GetTotalBooking implements booking.BookDataInterface.
func (repo *bookingQuery) GetTotalBooking() (int, error) {
	var count int64
	tx := repo.db.Model(&Booking{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

// GetTotalBookingPerYear implements booking.BookDataInterface.
func (repo *bookingQuery) GetTotalBookingPerMonth(year int, month int) (int, error) {
	var count int
	row := repo.db.Raw("SELECT COUNT(*) as count FROM bookings WHERE YEAR(created_at) = ? AND MONTH(created_at) = ?", year, month).Row()
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// PostBooking implements Booking.BookingDataInterface.
func (repo *bookingQuery) PostBooking(userId uint, input booking.BookingCore) (*booking.BookingCore, error) {
	var BookingGorm Booking

	boardingHouse := kd.BoardingHouse{}
	if err := repo.db.First(&boardingHouse, input.BoardingHouseId).Error; err != nil {
		return nil, err
	}
	var amount = boardingHouse.Price

	input.Total = float64(amount)
	input.UserID = userId

	payment, errPay := repo.paymentMidtrans.NewOrderPayment(input)
	if errPay != nil {
		return nil, errPay
	}

	fmt.Println(payment.ExpiredAt)
	repo.db.Transaction(func(tx *gorm.DB) error {
		BookingGorm = BookingCoreToModel(input)
		BookingGorm.PaymentType = payment.PaymentType
		BookingGorm.Status = payment.Status
		BookingGorm.VirtualNumber = payment.VirtualNumber
		BookingGorm.PaidAt = payment.PaidAt
		BookingGorm.ExpiredAt = payment.ExpiredAt
		BookingGorm.Total = float64(amount)
		if errBooking := tx.Create(&BookingGorm).Error; errBooking != nil {
			return errBooking
		}

		return nil
	})
	var BookingCores = ModelToCore(BookingGorm)

	return &BookingCores, nil
}

// GetBookings implements Booking.BookingDataInterface.
func (repo *bookingQuery) GetBookings(userId uint) ([]booking.BookingCore, error) {
	var BookingGorm []Booking
	tx := repo.db.Preload("BoardingHouse").Preload("BoardingHouse.Ratings").Preload("User").Find(&BookingGorm, "user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var kosRating []float64
	for _, booking := range BookingGorm {
		var ratings float64
		for _, rating := range booking.BoardingHouse.Ratings {
			ratings += float64(rating.Score)
			fmt.Println(rating.Score)
		}

		var resultRating float64
		if ratings > 0 {
			resultRating = ratings / float64(len(booking.BoardingHouse.Ratings))
		}
		kosRating = append(kosRating, resultRating)
	}
	fmt.Println("Kos Rating", kosRating)
	var BookingCores []booking.BookingCore
	for i, v := range BookingGorm {
		BookingCores = append(BookingCores, ModelToCore(v))

		BookingCores[i].Rating = kosRating[i]
		fmt.Println(i)
	}

	return BookingCores, nil
}

// CancelBooking implements Booking.BookingDataInterface.
func (repo *bookingQuery) CancelBooking(userId int, BookingId string, BookingCore booking.BookingCore) error {
	if BookingCore.Status == "cancelled" {
		repo.paymentMidtrans.CancelOrderPayment(BookingId)
	}

	dataGorm := Booking{
		Status: BookingCore.Status,
	}
	fmt.Println("Booking id::", BookingId)
	tx := repo.db.Model(&Booking{}).Where("id = ? AND user_id = ?", BookingId, userId).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// GetBooking implements Booking.BookingDataInterface.
func (repo *bookingQuery) GetBooking(userId uint) (*booking.BookingCore, error) {
	panic("unimplemented")
}

// WebhoocksData implements Booking.BookingDataInterface.
func (repo *bookingQuery) WebhoocksData(webhoocksReq booking.BookingCore) error {
	dataGorm := WebhoocksCoreToModel(webhoocksReq)
	tx := repo.db.Model(&Booking{}).Where("id = ?", webhoocksReq.ID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
