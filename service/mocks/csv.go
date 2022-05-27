// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	os "os"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// CSVMock is an autogenerated mock type for the CSV type
type CSVMock struct {
	mock.Mock
}

// MarshalFile provides a mock function with given fields: in, file
func (_m *CSVMock) MarshalFile(in interface{}, file *os.File) error {
	ret := _m.Called(in, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *os.File) error); ok {
		r0 = rf(in, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnmarshalFile provides a mock function with given fields: in, out
func (_m *CSVMock) UnmarshalFile(in *os.File, out interface{}) error {
	ret := _m.Called(in, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(*os.File, interface{}) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCSVMock creates a new instance of CSVMock. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSVMock(t testing.TB) *CSVMock {
	mock := &CSVMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
