package service

import (
	"employee-management-webapp/internal/repository"
	"employee-management-webapp/pkg/auth"
	"employee-management-webapp/pkg/hash"
	"employee-management-webapp/pkg/otp"
)

type EmployeesService struct {
	repo     repository.Employees
	hasher   hash.PasswordHasher
	tokenCfg auth.TokenConfig
	otpCfg   otp.Config

	// shiftsService Shi
}
