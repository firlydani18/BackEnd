package service

import (
	"KosKita/features/admin"
	"KosKita/features/user"
	"KosKita/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalData(t *testing.T) {
	repoUser := new(mocks.UserDataInterface)
	repoKos := new(mocks.KosDataInterface)
	repoBooking := new(mocks.BookDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repoUser, repoKos, repoBooking, srvUser)

	userData := &user.Core{
		ID:   1,
		Role: "admin",
	}

	dashboardData := admin.DashboardData{
		TotalUser:            90,
		TotalBooking:         88,
		TotalKos:             45,
		TotalBookingPerMonth: make([]int, 12),
	}

	t.Run("error from GetById", func(t *testing.T) {
		srvUser.On("GetById", 1).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{}, result)
		assert.EqualError(t, err, "database error")
		srvUser.AssertExpectations(t)
	})

	t.Run("error not admin", func(t *testing.T) {
		userData.Role = "user"
		srvUser.On("GetById", 1).Return(userData, nil).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{}, result)
		assert.EqualError(t, err, "anda bukan admin")
		srvUser.AssertExpectations(t)
	})

	t.Run("error from GetTotalUser", func(t *testing.T) {
		userData.Role = "admin"
		srvUser.On("GetById", 1).Return(userData, nil).Once()
		repoUser.On("GetTotalUser").Return(0, errors.New("database error")).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{}, result)
		assert.EqualError(t, err, "database error")
		repoUser.AssertExpectations(t)
	})

	t.Run("error from GetTotalKos", func(t *testing.T) {
		userData.Role = "admin"
		srvUser.On("GetById", 1).Return(userData, nil).Once()
		repoUser.On("GetTotalUser").Return(90, nil).Once()
		repoKos.On("GetTotalKos").Return(0, errors.New("database error")).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{TotalUser: 90}, result)
		assert.EqualError(t, err, "database error")
		repoKos.AssertExpectations(t)
	})

	t.Run("error from GetTotalBooking", func(t *testing.T) {
		userData.Role = "admin"
		srvUser.On("GetById", 1).Return(userData, nil).Once()
		repoUser.On("GetTotalUser").Return(90, nil).Once()
		repoKos.On("GetTotalKos").Return(45, nil).Once()
		repoBooking.On("GetTotalBooking").Return(0, errors.New("database error")).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{TotalUser: 90, TotalKos: 45}, result)
		assert.EqualError(t, err, "database error")
		repoBooking.AssertExpectations(t)
	})

	t.Run("error from GetTotalBookingPerMonth", func(t *testing.T) {
		userData.Role = "admin"
		srvUser.On("GetById", 1).Return(userData, nil).Once()
		repoUser.On("GetTotalUser").Return(90, nil).Once()
		repoKos.On("GetTotalKos").Return(45, nil).Once()
		repoBooking.On("GetTotalBooking").Return(88, nil).Once()
		repoBooking.On("GetTotalBookingPerMonth", 2024, 1).Return(0, errors.New("database error")).Once()

		result, err := srv.GetTotalData(1, 2024)

		assert.Error(t, err)
		assert.Equal(t, admin.DashboardData{TotalUser: 90, TotalKos: 45, TotalBooking: 88}, result)
		assert.EqualError(t, err, "database error")
		repoBooking.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		userData.Role = "admin"
		srvUser.On("GetById", 1).Return(userData, nil).Once()
		repoUser.On("GetTotalUser").Return(90, nil).Once()
		repoKos.On("GetTotalKos").Return(45, nil).Once()
		repoBooking.On("GetTotalBooking").Return(88, nil).Once()
		for month := 1; month <= 12; month++ {
			repoBooking.On("GetTotalBookingPerMonth", 2024, month).Return(0, nil).Once()
		}

		result, err := srv.GetTotalData(1, 2024)

		assert.NoError(t, err)
		assert.Equal(t, dashboardData, result)
		repoUser.AssertExpectations(t)
		repoKos.AssertExpectations(t)
		repoBooking.AssertExpectations(t)
		srvUser.AssertExpectations(t)
	})
}
