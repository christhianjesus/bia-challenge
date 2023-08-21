// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	period "github.com/christhianjesus/bia-challenge/internal/domain/period"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PeriodStrategy is an autogenerated mock type for the PeriodStrategy type
type PeriodStrategy struct {
	mock.Mock
}

// GeneratePeriods provides a mock function with given fields: startDate, endDate
func (_m *PeriodStrategy) GeneratePeriods(startDate time.Time, endDate time.Time) []period.Period {
	ret := _m.Called(startDate, endDate)

	var r0 []period.Period
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []period.Period); ok {
		r0 = rf(startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]period.Period)
		}
	}

	return r0
}

// NewPeriodStrategy creates a new instance of PeriodStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeriodStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *PeriodStrategy {
	mock := &PeriodStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}