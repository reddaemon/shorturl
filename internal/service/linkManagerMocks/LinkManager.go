// Code generated by mockery v2.14.0. DO NOT EDIT.

package linkManagerMocks

import mock "github.com/stretchr/testify/mock"

// LinkManager is an autogenerated mock type for the LinkManager type
type LinkManager struct {
	mock.Mock
}

// GetLink provides a mock function with given fields: shorturl
func (_m *LinkManager) GetLink(shorturl string) (string, error) {
	ret := _m.Called(shorturl)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(shorturl)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(shorturl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLink provides a mock function with given fields: shortlink, fullurl
func (_m *LinkManager) SetLink(shortlink string, fullurl string) (int64, error) {
	ret := _m.Called(shortlink, fullurl)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(shortlink, fullurl)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(shortlink, fullurl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLinkManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewLinkManager creates a new instance of LinkManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLinkManager(t mockConstructorTestingTNewLinkManager) *LinkManager {
	mock := &LinkManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}