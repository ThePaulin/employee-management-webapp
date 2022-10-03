package service

import (
	"context"
	"employee-management-webapp/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Managers

type ManagerLoginInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Managers interface {
	Login(ctx context.Context, input ManagerLoginInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, managerID primitive.ObjectID, hash string) error
	CreateWorkstation(ctx context.Context, managerID primitive.ObjectID, workstationName string, workstationCode string) (domain.Workstation, error)
}

type ConnectSendPulseInput struct {
	WorkstationID primitive.ObjectID
	ID            string
	Secret        string
	ListID        string
}

type Workstations interface {
	Create(ctx context.Context, name string) (primitive.ObjectID, error)
	GetByDomain(ctx context.Context, domainName string) (domain.Workstation, error)
	GetById(ctx context.Context, id primitive.ObjectID) (domain.Workstation, error)
	ConnectSendPulse(ctx context.Context, input ConnectSendPulseInput) error
}

// Employees
type EmployeeSignupInput struct {
	Firstname string
	Lastname  string
	Email     string
	Username  string
	Address   domain.Address
	Active    bool
}

type EmployeeSigninInput struct {
	Email    string
	Password string
}

type Employees interface {
	Signup(ctx context.Context, input EmployeeSignupInput) error
	Signin(ctx context.Context, input EmployeeSigninInput) error
	RefreshTokens(ctx context.Context, employeeId primitive.ObjectID, refreshToken string) (Tokens, error)
	GetById(ctx context.Context, employeeID primitive.ObjectID) (domain.Employee, error)
	GetByWorkstation(ctx context.Context, workstationID primitive.ObjectID, query domain.GetEmployeesQuery) (domain.Employees, int64, error)
	GiveAccessToSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error
	RemoveAccessToSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error
	GetBiweeklyWorkstations(ctx context.Context, employeeID primitive.ObjectID) domain.Employees
	GetBiweeklyShifts(ctx context.Context, employeeID primitive.ObjectID) domain.Employees
	GetBiweeklySchedules(ctx context.Context, employeeID primitive.ObjectID) domain.Employees
	GetSchedulesHistory(ctx context.Context, employeeID primitive.ObjectID) domain.Employees
	SetScheduleFinished(ctx context.Context, scheduleID primitive.ObjectID) error
}

// Schedules
type CreateScheduleInput struct {
	EmployeeID primitive.ObjectID
	ShiftID primitive.ObjectID
	WorkstationID primitive.ObjectID
	Date time.Time
}

type UpdateScheduleInput struct {
	ID string
	EmployeeID string
	ShiftID string
	WorkstationID string
}

type Schedules struct {
	Create(ctx context.Context, schedule CreateScheduleInput) (primitive.ObjectID, error)
	Update(ctx context.Context, scheduleID primitive.ObjectID, employeeID primitive.ObjectID) error
	Delete(ctx context.Context, scheduleID primitive.ObjectID) error
	GetById(ctx context.Context, scheduleID primitive.ObjectID) (domain.Schedule, error)
	GetByShift(ctx context.Context, shiftID primitive.ObjectID) (domain.Schedule, error)
	GetByIds(ctx context.Context, scheduleID []primitive.ObjectID) (domain.Schedules, error)
	GetAll(ctx context.Context, workstationID primitive.ObjectID) error
	GetByEmployee(ctx context.Context, employeeID primitive.ObjectID, query domain.GetSchedulesQuery) (domain.Schedules, int64, error)
}
