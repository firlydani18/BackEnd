package handler

import (
	"KosKita/features/booking"
	"time"
)

type BookingResponse struct {
	Id            string  `json:"booking_id"`
	UserID        uint    `json:"user_id"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	PaymentType   string  `json:"payment_method"`
	Total         float64 `json:"total"`
	Status        string  `json:"status"`
	Bank          string  `json:"bank"`
	VirtualNumber string  `json:"virtual_number"`
	ExpiredAt     string  `json:"expired_at"`
}

type BookingHistoryResponse struct {
	Id            string    `json:"order_id"`
	KosId         uint      `json:"kos_id"`
	KosName       string    `json:"kos_name"`
	KosFasilitas  []string  `json:"kos_fasilitas"`
	KosLokasi     string    `json:"kos_lokasi"`
	KosRating     float64   `json:"kos_rating"`
	StartDate     string    `json:"start_date"`
	KosMainFoto   string    `json:"kos_main_foto"`
	PaymentStatus string    `json:"payment_status"`
	TotalHarga    float64   `json:"total_harga"`
	CreatedAt     time.Time `json:"created_at"`
	PaidAt        string    `json:"paid_at"`
}

func CoreToResponse(o booking.BookingCore) BookingResponse {
	return BookingResponse{
		Id:            o.ID,
		UserID:        o.UserID,
		StartDate:     o.StartDate,
		EndDate:       o.EndDate,
		PaymentType:   o.PaymentType,
		Total:         o.Total,
		Status:        o.Status,
		Bank:          o.Bank,
		VirtualNumber: o.VirtualNumber,
		ExpiredAt:     o.ExpiredAt,
	}
}

func CoreToResponseBookingHistory(core booking.BookingCore) BookingHistoryResponse {
	return BookingHistoryResponse{
		Id:            core.ID,
		KosId:         core.BoardingHouse.ID,
		KosName:       core.BoardingHouse.Name,
		KosLokasi:     core.BoardingHouse.Address,
		StartDate:     core.StartDate,
		KosMainFoto:   core.BoardingHouse.PhotoMain,
		PaymentStatus: core.Status,
		TotalHarga:    core.Total,
		CreatedAt:     core.CreatedAt,
		PaidAt:        core.PaidAt,
		KosRating:     core.Rating,
	}
}
