package repository

import (
	"context"
	"employee-management-webapp/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employees interface {
	Create(ctx context.Context, employee *domain.Employee) error
	Update(ctx context.Context, inp domain.UpdateEmployeeInput) error
	Delete(ctx context.Context, employeeId primitive.ObjectID) error
	GetByCredentials(ctx context.Context, username string, password string)
	GetById(ctx context.Context, id primitive.ObjectID)
	GetBySchedules(ctx context.Context)
	GetByShift(ctx context.Context)
	AttachSchedule(ctx context.Context)
	DetachSchedule(ctx context.Context)
}
