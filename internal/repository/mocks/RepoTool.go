// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RepoTool is an autogenerated mock type for the RepoTool type
type RepoTool struct {
	mock.Mock
}

// Get provides a mock function with given fields: short
func (_m *RepoTool) Get(short string) (string, error) {
	ret := _m.Called(short)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(short)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(short)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: short, fullUrl
func (_m *RepoTool) Set(short string, fullUrl string) (int64, error) {
	ret := _m.Called(short, fullUrl)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(short, fullUrl)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(short, fullUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepoTool interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoTool creates a new instance of RepoTool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoTool(t mockConstructorTestingTNewRepoTool) *RepoTool {
	mock := &RepoTool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}