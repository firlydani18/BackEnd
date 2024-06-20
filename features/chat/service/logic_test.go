package service

import (
	"KosKita/features/chat"
	"KosKita/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoom(t *testing.T) {
	repo := new(mocks.ChatDataInterface)
	srv := New(repo)

	returnData := chat.Core{
		RoomID:     "87677",
		ReceiverID: 2,
	}

	t.Run("error from repo", func(t *testing.T) {
		repo.On("GetRoom", 1).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetRoom(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("GetRoom", 1).Return([]chat.Core{returnData}, nil).Once()

		result, err := srv.GetRoom(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result[0])

		repo.AssertExpectations(t)
	})
}

func TestCreateRoom(t *testing.T) {
	repo := new(mocks.ChatDataInterface)
	srv := New(repo)

	t.Run("success", func(t *testing.T) {
		repo.On("CreateRoom", "room1", 1, 2).Return(nil).Once()

		err := srv.CreateRoom("room1", 1, 2)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreateChat(t *testing.T) {
	repo := new(mocks.ChatDataInterface)
	srv := New(repo)

	inputChat := chat.Core{
		Message:    "Hello",
		RoomID:     "room1",
		ReceiverID: 2,
		SenderID:   1,
	}

	t.Run("error from repo", func(t *testing.T) {
		repo.On("CreateMessage", 2, 1, inputChat).Return(chat.Core{}, errors.New("database error")).Once()

		result, err := srv.CreateChat(2, 1, inputChat)

		assert.Error(t, err)
		assert.Equal(t, chat.Core{}, result)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("CreateMessage", 2, 1, inputChat).Return(inputChat, nil).Once()

		result, err := srv.CreateChat(2, 1, inputChat)

		assert.NoError(t, err)
		assert.Equal(t, inputChat, result)
		repo.AssertExpectations(t)
	})
}

func TestGetMessage(t *testing.T) {
	repo := new(mocks.ChatDataInterface)
	srv := New(repo)

	returnData := []chat.Core{
		{
			Message:    "Hello",
			RoomID:     "room1",
			ReceiverID: 2,
			SenderID:   1,
		},
	}

	t.Run("error from repo", func(t *testing.T) {
		repo.On("GetMessage", "room1").Return(nil, errors.New("database error")).Once()

		result, err := srv.GetMessage("room1")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("GetMessage", "room1").Return(returnData, nil).Once()

		result, err := srv.GetMessage("room1")

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}