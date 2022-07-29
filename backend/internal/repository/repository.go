package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Managers interface {
	Create(ctx context.Context, manager *domain.Manager) error
	GetByCredentials(ctx context.Context, email string, password string) (domain.Manager, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Manager, error)
	SetSession(ctx context.Context, managerID primitive.ObjectID, session domain.Session) error
	AttachWorkstation(ctx context.Context, managerID primitive.ObjectID, workstationID primitive.ObjectID) error
}

type UpdateWorkstationInput struct {
	Name           *string
	Description    *string
	Status         *string
	Managers       []primitive.ObjectID
	BiweeklyShifts []primitive.ObjectID
	ShiftsHistory  []primitive.ObjectID
}

type Workstations interface {
	Create(ctx context.Context, code int, name string) (primitive.ObjectID, error)
	GetById(ctx context.Context, workstationID primitive.ObjectID) (domain.Workstation, error)
	GetByManager(ctx context.Context, managerID primitive.ObjectID) (domain.Workstations, error)
	GetByShift(ctx context.Context, shiftID primitive.ObjectID) (domain.Workstation, error)
	Update(ctx context.Context, workstationID primitive.ObjectID, input UpdateWorkstationInput) error
}

type Employees interface {
	Create(ctx context.Context, employee *domain.Employee) error
	Update(ctx context.Context, employee domain.UpdateEmployeeInput) error
	Delete(ctx context.Context, employeeID primitive.ObjectID) error
	GetByCredentials(ctx context.Context, username string, password string) (domain.Employee, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Employee, error)
	GetById(ctx context.Context, employeeID primitive.ObjectID) (domain.Employee, error)
	GetByWorkstation(ctx context.Context, workstationID primitive.ObjectID, query domain.GetEmployeesQuery) (domain.Employees, int64, error)
	SetSession(ctx context.Context, employeeID primitive.ObjectID, session domain.Session) error
	AttachSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error
	DetachSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error
}

type UpdateShiftInput struct {
	Type    *string
	StartAt *time.Time
	EndAt   *time.Time
	Status  *string
}

type Shifts interface {
	Create(ctx context.Context, shift domain.Shift) (primitive.ObjectID, error)
	Update(ctx context.Context, input UpdateShiftInput) error
	Delete(ctx context.Context, shiftID primitive.ObjectID) error
	GetById(ctx context.Context, shiftID primitive.ObjectID) (domain.Shift, error)
	GetByStatus(ctx context.Context, shiftID primitive.ObjectID) (domain.Shifts, error)
	SetStatus(ctx context.Context, shiftID primitive.ObjectID, status string) error
}

type UpdateScheduleInput struct {
	WorkstationID primitive.ObjectID
	EmployeeID    primitive.ObjectID
	ShiftID       primitive.ObjectID
}

type Schedules interface {
	Create(ctx context.Context, schedule domain.Schedule) (primitive.ObjectID, error)
	Update(ctx context.Context, scheduleID primitive.ObjectID, input UpdateScheduleInput) error
	Delete(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID) error
	GetByEmployee(ctx context.Context, employeeID primitive.ObjectID) (domain.Schedules, error)
	GetById(ctx context.Context, scheduleID primitive.ObjectID) (domain.Schedule, error)
	GetByShift(ctx context.Context, shiftID primitive.ObjectID) (domain.Schedule, error)
	GetByIds(ctx context.Context, scheduleIDs []primitive.ObjectID) (domain.Schedules, error)
}

type Repositories struct {
	Managers     Managers
	Workstations Workstations
	Employees    Employees
	Shifts       Shifts
	Schedules    Schedules
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Managers:     NewManagersRepo(db),
		Workstations: NewWorkstationsRepo(db),
		Employees:    NewEmployeesRepo(db),
		// Shifts:       NewShiftsRepo(db),
		// Schedules:    NewSchedulesRepo(db),
	}
}

func getPaginationOpts(pagination *domain.PaginationQuery) *options.FindOptions {
	var opts *options.FindOptions
	if pagination != nil {
		opts = &options.FindOptions{
			Skip:  pagination.GetSkip(),
			Limit: pagination.GetLimit(),
		}
	}
	return opts
}

func filterDateQueries(dateFrom string, dateTo string, fieldName string, filter bson.M) error {
	if dateFrom != "" && dateTo != "" {
		dateFrom, err := time.Parse(time.RFC3339, dateFrom)
		if err != nil {
			return err
		}

		dateTo, err := time.Parse(time.RFC3339, dateTo)
		if err != nil {
			return err
		}

		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			"$and": []bson.M{
				{fieldName: bson.M{"$gte": dateFrom}},
				{fieldName: bson.M{"$lte": dateTo}},
			},
		})
	}

	if dateFrom != "" && dateTo == "" {
		dateFrom, err := time.Parse(time.RFC3339, dateFrom)
		if err != nil {
			return err
		}

		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			fieldName: bson.M{"$gte": dateFrom},
		})
	}

	if dateFrom == "" && dateTo != "" {
		dateTo, err := time.Parse(time.RFC3339, dateTo)
		if err != nil {
			return err
		}

		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			fieldName: bson.M{"$lte": dateTo},
		})
	}

	return nil
}
