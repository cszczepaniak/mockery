// Code generated by mockery. DO NOT EDIT.

package example_project

import mock "github.com/stretchr/testify/mock"

// MockStringer is an autogenerated mock type for the Stringer type
type MockStringer struct {
	mock.Mock
}

type MockStringer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStringer) EXPECT() *MockStringer_Expecter {
	return &MockStringer_Expecter{mock: &_m.Mock}
}

// String provides a mock function with given fields:
func (_m *MockStringer) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStringer_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockStringer_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockStringer_Expecter) String() *MockStringer_String_Call {
	return &MockStringer_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockStringer_String_Call) Run(run func()) *MockStringer_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStringer_String_Call) Return(_a0 string) *MockStringer_String_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockStringer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStringer creates a new instance of MockStringer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStringer(t mockConstructorTestingTNewMockStringer) *MockStringer {
	mock := &MockStringer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
