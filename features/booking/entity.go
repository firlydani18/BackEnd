package booking

import (
	kd "KosKita/features/kos"
	ud "KosKita/features/user"
	"time"
)

type BookingCore struct {
	ID              string
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
	Rating          float64
	CreatedAt       time.Time
	User            ud.Core
	BoardingHouse   kd.Core
}

type BookDataInterface interface {
	PostBooking(userId uint, input BookingCore) (*BookingCore, error)
	GetBooking(userId uint) (*BookingCore, error)
	GetBookings(userId uint) ([]BookingCore, error)
	CancelBooking(userId int, bookingId string, bookingCore BookingCore) error
	WebhoocksData(webhoocksReq BookingCore) error
	GetTotalBooking() (int, error)
	GetTotalBookingPerMonth(year int, month int) (int, error)
}

type BookingServiceInterface interface {
	PostBooking(userId uint, input BookingCore) (*BookingCore, error)
	GetBookings(userId uint) ([]BookingCore, error)
	CancelBooking(userId int, bookingId string, bookingCore BookingCore) error
	WebhoocksService(webhoocksReq BookingCore) error
}
