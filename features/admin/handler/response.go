package handler

import "KosKita/features/admin"

type DashboardData struct {
	TotalUser            int            `json:"total_user"`
	TotalBooking         int            `json:"total_booking"`
	TotalKos             int            `json:"total_kos"`
	TotalBookingPerMonth map[string]int `json:"total_booking_per_month"`
}

func CoreToResponseDashboard(core *admin.DashboardData) DashboardData {
	bulan := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	totalBookingPerMonth := make(map[string]int)
	for i, jumlah := range core.TotalBookingPerMonth {
		totalBookingPerMonth[bulan[i]] = jumlah
	}

	return DashboardData{
		TotalUser:            core.TotalUser,
		TotalBooking:         core.TotalBooking,
		TotalKos:             core.TotalKos,
		TotalBookingPerMonth: totalBookingPerMonth,
	}
}
