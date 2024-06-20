// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	admin "KosKita/features/admin"

	mock "github.com/stretchr/testify/mock"
)

// AdminServiceInterface is an autogenerated mock type for the AdminServiceInterface type
type AdminServiceInterface struct {
	mock.Mock
}

// GetTotalData provides a mock function with given fields: userIdLogin, uear
func (_m *AdminServiceInterface) GetTotalData(userIdLogin int, uear int) (admin.DashboardData, error) {
	ret := _m.Called(userIdLogin, uear)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalData")
	}

	var r0 admin.DashboardData
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (admin.DashboardData, error)); ok {
		return rf(userIdLogin, uear)
	}
	if rf, ok := ret.Get(0).(func(int, int) admin.DashboardData); ok {
		r0 = rf(userIdLogin, uear)
	} else {
		r0 = ret.Get(0).(admin.DashboardData)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userIdLogin, uear)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAdminServiceInterface creates a new instance of AdminServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminServiceInterface {
	mock := &AdminServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}