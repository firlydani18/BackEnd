package service

import (
	"KosKita/features/booking"
	"errors"
)

type bookingService struct {
	bookingData booking.BookDataInterface
}

func NewBooking(repo booking.BookDataInterface) booking.BookingServiceInterface {
	return &bookingService{
		bookingData: repo,
	}
}

// PostBooking implements Booking.BookingServiceInterface.
func (service *bookingService) PostBooking(userId uint, input booking.BookingCore) (*booking.BookingCore, error) {
	if userId <= 0 {
		return nil, errors.New("invalid id")
	}
	res, err := service.bookingData.PostBooking(userId, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetBookings implements Booking.BookingServiceInterface.
func (service *bookingService) GetBookings(userId uint) ([]booking.BookingCore, error) {
	results, err := service.bookingData.GetBookings(userId)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CancelBooking implements Booking.BookingServiceInterface.
func (service *bookingService) CancelBooking(userId int, bookingId string, bookingCore booking.BookingCore) error {
	if bookingCore.Status == "" {
		bookingCore.Status = "cancelled"
	}
	err := service.bookingData.CancelBooking(userId, bookingId, bookingCore)
	return err
}

// WebhoocksService implements Booking.BookingServiceInterface.
func (service *bookingService) WebhoocksService(webhoocksReq booking.BookingCore) error {
	if webhoocksReq.ID == "" {
		return errors.New("invalid Booking id")
	}

	err := service.bookingData.WebhoocksData(webhoocksReq)
	if err != nil {
		return err
	}

	return nil
}
