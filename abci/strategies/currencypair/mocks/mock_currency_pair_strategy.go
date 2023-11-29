// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CurrencyPairStrategy is an autogenerated mock type for the CurrencyPairStrategy type
type CurrencyPairStrategy struct {
	mock.Mock
}

// FromID provides a mock function with given fields: ctx, id
func (_m *CurrencyPairStrategy) FromID(ctx types.Context, id uint64) (oracletypes.CurrencyPair, error) {
	ret := _m.Called(ctx, id)

	var r0 oracletypes.CurrencyPair
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (oracletypes.CurrencyPair, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) oracletypes.CurrencyPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(oracletypes.CurrencyPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields: ctx, cp
func (_m *CurrencyPairStrategy) ID(ctx types.Context, cp oracletypes.CurrencyPair) (uint64, error) {
	ret := _m.Called(ctx, cp)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair) (uint64, error)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair) uint64); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context, oracletypes.CurrencyPair) error); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCurrencyPairStrategy creates a new instance of CurrencyPairStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCurrencyPairStrategy(t interface {
	mock.TestingT
	Cleanup(func())
},
) *CurrencyPairStrategy {
	mock := &CurrencyPairStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}