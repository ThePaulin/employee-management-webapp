package otp

import "github.com/stretchr/testify/mock"

type MockConfig struct {
	mock.Mock
}

func (mock *MockConfig) RandomSecret(length int) string {
	args := mock.Called(length)

	return args.Get(0).(string)
}
