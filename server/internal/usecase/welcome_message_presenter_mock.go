// Code generated by mockery v2.53.0. DO NOT EDIT.

package usecase

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// MockWelcomeMessagePresenter is an autogenerated mock type for the WelcomeMessagePresenter type
type MockWelcomeMessagePresenter struct {
	mock.Mock
}

type MockWelcomeMessagePresenter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWelcomeMessagePresenter) EXPECT() *MockWelcomeMessagePresenter_Expecter {
	return &MockWelcomeMessagePresenter_Expecter{mock: &_m.Mock}
}

// Render provides a mock function with given fields: message, userID
func (_m *MockWelcomeMessagePresenter) Render(message string, userID string) (string, *model.AppError) {
	ret := _m.Called(message, userID)

	if len(ret) == 0 {
		panic("no return value specified for Render")
	}

	var r0 string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) (string, *model.AppError)); ok {
		return rf(message, userID)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(message, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(message, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockWelcomeMessagePresenter_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type MockWelcomeMessagePresenter_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - message string
//   - userID string
func (_e *MockWelcomeMessagePresenter_Expecter) Render(message interface{}, userID interface{}) *MockWelcomeMessagePresenter_Render_Call {
	return &MockWelcomeMessagePresenter_Render_Call{Call: _e.mock.On("Render", message, userID)}
}

func (_c *MockWelcomeMessagePresenter_Render_Call) Run(run func(message string, userID string)) *MockWelcomeMessagePresenter_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockWelcomeMessagePresenter_Render_Call) Return(_a0 string, _a1 *model.AppError) *MockWelcomeMessagePresenter_Render_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWelcomeMessagePresenter_Render_Call) RunAndReturn(run func(string, string) (string, *model.AppError)) *MockWelcomeMessagePresenter_Render_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWelcomeMessagePresenter creates a new instance of MockWelcomeMessagePresenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWelcomeMessagePresenter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWelcomeMessagePresenter {
	mock := &MockWelcomeMessagePresenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
