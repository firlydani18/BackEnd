package handler

import (
	"KosKita/features/booking"
	"fmt"
	"math/rand"
	"time"
	// "github.com/google/uuid"
)

type BookingRequest struct {
	ID              string
	BoardingHouseId uint   `json:"kos_id"`
	Method          string `json:"payment_type" form:"payment_type"`
	Bank            string `json:"bank" form:"bank"`
	StartDate       string `json:"start_date" form:"start_date"`
	EndDate         string `json:"end_date" form:"end_date"`
}

type CancelOrderRequest struct {
	Status string `json:"status"`
}

type CancelBookingRequest struct {
	Status string `json:"status"`
}

type WebhoocksRequest struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	SignatureKey      string `json:"signature_key"`
	TransactionTime   string `json:"transaction_time"`
}

func RequestToCoreBooking(input BookingRequest) booking.BookingCore {
	return booking.BookingCore{
		ID:              generateRoomID(),
		BoardingHouseId: input.BoardingHouseId,
		StartDate:       input.StartDate,
		EndDate:         input.EndDate,
		PaymentType:     input.Method,
		Bank:            input.Bank,
	}
}

func CancelRequestToCoreBooking(input CancelBookingRequest) booking.BookingCore {
	return booking.BookingCore{
		Status: input.Status,
	}
}

func WebhoocksRequestToCore(input WebhoocksRequest) booking.BookingCore {
	return booking.BookingCore{
		ID:     input.OrderID,
		Status: input.TransactionStatus,
		PaidAt: input.TransactionTime,
	}
}

func generateRoomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%05d", rand.Intn(100000))
}
