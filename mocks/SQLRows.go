// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SQLRows is an autogenerated mock type for the SQLRows type
type SQLRows struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *SQLRows) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *SQLRows) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *SQLRows) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *SQLRows) Scan(dest ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dest...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSQLRows creates a new instance of SQLRows. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSQLRows(t interface {
	mock.TestingT
	Cleanup(func())
}) *SQLRows {
	mock := &SQLRows{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
