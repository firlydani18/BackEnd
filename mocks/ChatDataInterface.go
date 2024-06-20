// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	chat "KosKita/features/chat"

	mock "github.com/stretchr/testify/mock"
)

// ChatDataInterface is an autogenerated mock type for the ChatDataInterface type
type ChatDataInterface struct {
	mock.Mock
}

// CreateMessage provides a mock function with given fields: receiverID, senderID, input
func (_m *ChatDataInterface) CreateMessage(receiverID int, senderID int, input chat.Core) (chat.Core, error) {
	ret := _m.Called(receiverID, senderID, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 chat.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, chat.Core) (chat.Core, error)); ok {
		return rf(receiverID, senderID, input)
	}
	if rf, ok := ret.Get(0).(func(int, int, chat.Core) chat.Core); ok {
		r0 = rf(receiverID, senderID, input)
	} else {
		r0 = ret.Get(0).(chat.Core)
	}

	if rf, ok := ret.Get(1).(func(int, int, chat.Core) error); ok {
		r1 = rf(receiverID, senderID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoom provides a mock function with given fields: roomID, receiverID, senderID
func (_m *ChatDataInterface) CreateRoom(roomID string, receiverID int, senderID int) error {
	ret := _m.Called(roomID, receiverID, senderID)

	if len(ret) == 0 {
		panic("no return value specified for CreateRoom")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, int) error); ok {
		r0 = rf(roomID, receiverID, senderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessage provides a mock function with given fields: roomId
func (_m *ChatDataInterface) GetMessage(roomId string) ([]chat.Core, error) {
	ret := _m.Called(roomId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessage")
	}

	var r0 []chat.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]chat.Core, error)); ok {
		return rf(roomId)
	}
	if rf, ok := ret.Get(0).(func(string) []chat.Core); ok {
		r0 = rf(roomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chat.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoom provides a mock function with given fields: userIdlogin
func (_m *ChatDataInterface) GetRoom(userIdlogin int) ([]chat.Core, error) {
	ret := _m.Called(userIdlogin)

	if len(ret) == 0 {
		panic("no return value specified for GetRoom")
	}

	var r0 []chat.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]chat.Core, error)); ok {
		return rf(userIdlogin)
	}
	if rf, ok := ret.Get(0).(func(int) []chat.Core); ok {
		r0 = rf(userIdlogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chat.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userIdlogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatDataInterface creates a new instance of ChatDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatDataInterface {
	mock := &ChatDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}