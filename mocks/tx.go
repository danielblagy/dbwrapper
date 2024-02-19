// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dbwrapper "github.com/danielblagy/dbwrapper"
	mock "github.com/stretchr/testify/mock"
)

// Tx is an autogenerated mock type for the Tx type
type Tx struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *Tx) BeginTx(ctx context.Context) (dbwrapper.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 dbwrapper.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (dbwrapper.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) dbwrapper.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbwrapper.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: ctx
func (_m *Tx) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *Tx) Exec(ctx context.Context, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, dest, query, args
func (_m *Tx) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSingle provides a mock function with given fields: ctx, dest, query, args
func (_m *Tx) GetSingle(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSingle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx
func (_m *Tx) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTx creates a new instance of Tx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tx {
	mock := &Tx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
