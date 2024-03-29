package domain

import "errors"

var (
	ErrManagerNotFound          = errors.New("manager doesn't exist")
	ErrManagerAlreadyExists     = errors.New("manager with such email or phone number already exists")
	ErrWorkstationAlreadyExists = errors.New("workstation with such name or code already exists")
	ErrShiftAlreadyExists       = errors.New("shift with such code already exists")
	ErrScheduleNotFound         = errors.New("schedule doesn't exist")
	ErrShiftNotFound            = errors.New("shift doesn't exist")
	ErrWorkstationNotFound      = errors.New("workstation doesn't exist")
	ErrEmployeeNotFound         = errors.New("employee doesn't exist")
	ErrWorkstationtNotAvailable = errors.New("workstation is not available")
	ErrScheduleInvalid          = errors.New("schedule is not valid")
	ErrEmployeeAlreadyExists    = errors.New("employee with such email or phone number already exists")
	ErrEmployeeInactive         = errors.New("employee is deactivated by the managers")
)
