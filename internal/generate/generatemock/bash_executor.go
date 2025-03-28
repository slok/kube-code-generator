// Code generated by mockery v2.53.3. DO NOT EDIT.

package generatemock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BashExecutor is an autogenerated mock type for the BashExecutor type
type BashExecutor struct {
	mock.Mock
}

// BashExec provides a mock function with given fields: ctx, bashCmd
func (_m *BashExecutor) BashExec(ctx context.Context, bashCmd string) (string, error) {
	ret := _m.Called(ctx, bashCmd)

	if len(ret) == 0 {
		panic("no return value specified for BashExec")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, bashCmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, bashCmd)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bashCmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBashExecutor creates a new instance of BashExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBashExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *BashExecutor {
	mock := &BashExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
