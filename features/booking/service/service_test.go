package service

import (
	"KosKita/features/booking"
	"KosKita/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBooking(t *testing.T) {
	repo := new(mocks.BookDataInterface)
	srv := NewBooking(repo)

	input := booking.BookingCore{
		ID:              "10843",
		UserID:          1,
		BoardingHouseId: 1,
		StartDate:       "2024-01-01",
		EndDate:         "2024-12-31",
		PaymentType:     "credit_card",
		Total:           100000,
		Bank:            "BCA",
	}

	t.Run("invalid user id", func(t *testing.T) {
		repo.On("PostBooking", uint(0), input).Return(nil, errors.New("invalid id")).Once()

		_, err := srv.PostBooking(0, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("PostBooking", uint(1), input).Return(nil, errors.New("database error")).Once()

		_, err := srv.PostBooking(1, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("PostBooking", uint(1), input).Return(&input, nil).Once()

		result, err := srv.PostBooking(1, input)

		assert.NoError(t, err)
		assert.Equal(t, input, *result)
	})
}

func TestGetBookings(t *testing.T) {
	repo := new(mocks.BookDataInterface)
	srv := NewBooking(repo)

	returnBook := []booking.BookingCore{
		{
			ID:              "10843",
			UserID:          1,
			BoardingHouseId: 1,
			StartDate:       "2024-01-01",
			EndDate:         "2024-12-31",
			PaymentType:     "credit_card",
			Total:           100000,
			Bank:            "BCA",
		},
	}

	t.Run("invalid user id", func(t *testing.T) {
		repo.On("GetBookings", uint(0)).Return(nil, errors.New("invalid id")).Once()

		_, err := srv.GetBookings(0)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("GetBookings", uint(1)).Return(nil, errors.New("database error")).Once()

		_, err := srv.GetBookings(1)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {

		repo.On("GetBookings", uint(1)).Return(returnBook, nil).Once()

		result, err := srv.GetBookings(1)

		assert.NoError(t, err)
		assert.Equal(t, returnBook, result)
	})
}

func TestCancelBooking(t *testing.T) {
	repo := new(mocks.BookDataInterface)
	srv := NewBooking(repo)

	t.Run("invalid user id", func(t *testing.T) {
		repo.On("CancelBooking", 0, "10843", booking.BookingCore{Status: "cancelled"}).Return(errors.New("invalid id")).Once()

		err := srv.CancelBooking(0, "10843", booking.BookingCore{})

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("CancelBooking", 1, "10843", booking.BookingCore{Status: "cancelled"}).Return(errors.New("database error")).Once()

		err := srv.CancelBooking(1, "10843", booking.BookingCore{})

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("CancelBooking", 1, "10843", booking.BookingCore{Status: "cancelled"}).Return(nil).Once()

		err := srv.CancelBooking(1, "10843", booking.BookingCore{})

		assert.NoError(t, err)
	})
}

func TestWebhoocksService(t *testing.T) {
	repo := new(mocks.BookDataInterface)
	srv := NewBooking(repo)

	t.Run("invalid booking id", func(t *testing.T) {
		err := srv.WebhoocksService(booking.BookingCore{ID: ""})

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid Booking id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("WebhoocksData", booking.BookingCore{ID: "10843", Status: "cancelled"}).Return(errors.New("database error")).Once()

		err := srv.WebhoocksService(booking.BookingCore{ID: "10843", Status: "cancelled"})

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("WebhoocksData", booking.BookingCore{ID: "10843", Status: "cancelled"}).Return(nil).Once()

		err := srv.WebhoocksService(booking.BookingCore{ID: "10843", Status: "cancelled"})

		assert.NoError(t, err)
	})
}
