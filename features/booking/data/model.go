package data

import (
	"KosKita/features/booking"
	"KosKita/features/kos"
	kd "KosKita/features/kos/data"
	"KosKita/features/user"
	ud "KosKita/features/user/data"

	"gorm.io/gorm"
)

type Booking struct {
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`
	gorm.Model
	UserID          uint
	BoardingHouseId uint
	StartDate       string
	EndDate         string
	PaymentType     string
	Total           float64
	Status          string
	Bank            string
	VirtualNumber   string
	ExpiredAt       string
	PaidAt          string
	User            ud.User
	BoardingHouse   kd.BoardingHouse
}

func BookingCoreToModel(o booking.BookingCore) Booking {
	return Booking{
		ID:              o.ID,
		UserID:          o.UserID,
		BoardingHouseId: o.BoardingHouseId,
		StartDate:       o.StartDate,
		EndDate:         o.EndDate,
		PaymentType:     o.PaymentType,
		Total:           o.Total,
		Status:          o.Status,
		Bank:            o.Bank,
		VirtualNumber:   o.VirtualNumber,
		ExpiredAt:       o.ExpiredAt,
		PaidAt:          o.PaidAt,
	}
}

func ModelToCore(o Booking) booking.BookingCore {
	return booking.BookingCore{
		ID:              o.ID,
		UserID:          o.UserID,
		BoardingHouseId: o.BoardingHouseId,
		StartDate:       o.StartDate,
		EndDate:         o.EndDate,
		PaymentType:     o.PaymentType,
		Total:           o.Total,
		Status:          o.Status,
		Bank:            o.Bank,
		VirtualNumber:   o.VirtualNumber,
		ExpiredAt:       o.ExpiredAt,
		PaidAt:          o.PaidAt,
		CreatedAt:       o.CreatedAt,
		User: user.Core{
			ID:           o.User.ID,
			Name:         o.User.Name,
			UserName:     o.User.UserName,
			Email:        o.User.Email,
			Password:     o.User.Password,
			Gender:       o.User.Gender,
			Role:         o.User.Role,
			PhotoProfile: o.User.PhotoProfile,
		},
		BoardingHouse: kos.Core{
			ID:          o.BoardingHouse.ID,
			Name:        o.BoardingHouse.Name,
			Description: o.BoardingHouse.Description,
			Category:    o.BoardingHouse.Category,
			Price:       o.BoardingHouse.Price,
			Rooms:       o.BoardingHouse.Rooms,
			Address:     o.BoardingHouse.Address,
			Longitude:   o.BoardingHouse.Longitude,
			Latitude:    o.BoardingHouse.Latitude,
			PhotoMain:   o.BoardingHouse.PhotoMain,
			UserID:      o.BoardingHouse.UserID,
			Ratings:     []kos.RatingCore{},
		},
	}
}

func WebhoocksCoreToModel(reqNotif booking.BookingCore) Booking {
	return Booking{
		Status: reqNotif.Status,
		PaidAt: reqNotif.PaidAt,
	}
}
