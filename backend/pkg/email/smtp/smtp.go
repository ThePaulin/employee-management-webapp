package smtp

import (
	"employee-management-webapp/pkg/email"
	"errors"
	"fmt"

	"github.com/go-gomail/gomail"
)

type SMTPSender struct {
	from string
	pass string
	host string
	port int
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	if !email.IsEmailValid(from) {
		return nil, errors.New("invalid `from email` field")
	}

	return &SMTPSender{from: from, pass: pass, host: host, port: port}, nil
}

func (smtp *SMTPSender) Send(input email.SendEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", smtp.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)

	dialer := gomail.NewDialer(smtp.host, smtp.port, smtp.from, smtp.pass)
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email via smtp: %s", err)
	}

	return nil
}
