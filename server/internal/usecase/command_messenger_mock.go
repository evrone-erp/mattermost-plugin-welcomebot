// Code generated by mockery v2.53.0. DO NOT EDIT.

package usecase

import mock "github.com/stretchr/testify/mock"

// MockCommandMessenger is an autogenerated mock type for the CommandMessenger type
type MockCommandMessenger struct {
	mock.Mock
}

type MockCommandMessenger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommandMessenger) EXPECT() *MockCommandMessenger_Expecter {
	return &MockCommandMessenger_Expecter{mock: &_m.Mock}
}

// PostCommandResponse provides a mock function with given fields: message
func (_m *MockCommandMessenger) PostCommandResponse(message string) {
	_m.Called(message)
}

// MockCommandMessenger_PostCommandResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCommandResponse'
type MockCommandMessenger_PostCommandResponse_Call struct {
	*mock.Call
}

// PostCommandResponse is a helper method to define mock.On call
//   - message string
func (_e *MockCommandMessenger_Expecter) PostCommandResponse(message interface{}) *MockCommandMessenger_PostCommandResponse_Call {
	return &MockCommandMessenger_PostCommandResponse_Call{Call: _e.mock.On("PostCommandResponse", message)}
}

func (_c *MockCommandMessenger_PostCommandResponse_Call) Run(run func(message string)) *MockCommandMessenger_PostCommandResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCommandMessenger_PostCommandResponse_Call) Return() *MockCommandMessenger_PostCommandResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCommandMessenger_PostCommandResponse_Call) RunAndReturn(run func(string)) *MockCommandMessenger_PostCommandResponse_Call {
	_c.Run(run)
	return _c
}

// NewMockCommandMessenger creates a new instance of MockCommandMessenger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommandMessenger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommandMessenger {
	mock := &MockCommandMessenger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
