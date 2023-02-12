// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	"unsafe"

	mock "github.com/stretchr/testify/mock"
)

// UnsafeInterface is an autogenerated mock type for the UnsafeInterface type
type UnsafeInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: ptr
func (_m *UnsafeInterface) Do(ptr *unsafe.Pointer) {
	_m.Called(ptr)
}

type mockConstructorTestingTNewUnsafeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeInterface creates a new instance of UnsafeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeInterface(t mockConstructorTestingTNewUnsafeInterface) *UnsafeInterface {
	mock := &UnsafeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
