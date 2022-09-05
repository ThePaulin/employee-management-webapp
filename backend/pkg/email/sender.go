package email

import (
	"bytes"
	"employee-management-webapp/pkg/logger"
	"errors"
	"html/template"
)

type SendEmailInput struct {
	To      string
	Subject string
	Body    string
}

type Sender interface {
	Send(input SendEmailInput) error
}

func (email *SendEmailInput) GenerateBodyFromHTML(templateFileName string, data interface{}) error {
	tpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		logger.Error("failed to parse file %s:%s", templateFileName, err.Error())

		return err
	}

	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, data); err != nil {
		return err
	}

	email.Body = buf.String()

	return nil
}

func (email *SendEmailInput) Validate() error {
	if email.To == "" {
		return errors.New("email `to` not set")
	}

	if email.Subject == "" || email.Body == "" {
		return errors.New("subject/body not set")
	}

	if !IsEmailValid(email.To) {
		return errors.New("email not valid")
	}

	return nil
}
