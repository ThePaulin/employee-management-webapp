package mock_email

import (
	"employee-management-webapp/pkg/email"

	"github.com/stretchr/testify/mock"
)

type EmailProvider struct {
	mock.Mock
}

type EmailSender struct {
	mock.Mock
}

func (provider *EmailProvider) AddEmailToList(input email.AddEmailInput) error {
	args := provider.Called(input)

	return args.Error(0)
}

func (sender *EmailSender) Send(input email.SendEmailInput) error {
	args := sender.Called(input)

	return args.Error(0)
}
