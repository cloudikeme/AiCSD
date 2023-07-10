// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "aicsd/pkg/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// MatchTask provides a mock function with given fields: entry
func (_m *Client) MatchTask(entry types.Job) (bool, error) {
	ret := _m.Called(entry)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Job) (bool, error)); ok {
		return rf(entry)
	}
	if rf, ok := ret.Get(0).(func(types.Job) bool); ok {
		r0 = rf(entry)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Job) error); ok {
		r1 = rf(entry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveById provides a mock function with given fields: id
func (_m *Client) RetrieveById(id string) (types.Task, error) {
	ret := _m.Called(id)

	var r0 types.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) types.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(types.Task)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}