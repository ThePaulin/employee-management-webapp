package otp

import "github.com/xlzd/gotp"

type Config interface {
	RandomSecret(lenght int) string
}

type OTPConfig struct{}

func NewOTPConfig() *OTPConfig {
	return &OTPConfig{}
}

func (config *OTPConfig) RandomSecret(length int) string {
	return gotp.RandomSecret(length)
}
