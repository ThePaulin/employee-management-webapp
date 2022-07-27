package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"employee-management-webapp/pkg/database/mongodb"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeesRepo struct {
	db *mongo.Collection
}

func NewEmployeesRepo(db *mongo.Database) *EmployeesRepo {
	return &EmployeesRepo{
		db: db.Collection(employeesCollection),
	}
}

func (r *EmployeesRepo) Create(ctx context.Context, employee *domain.Employee) error {
	res, err := r.db.InsertOne(ctx, employee)
	if mongodb.IsDuplicate(err) {
		return domain.ErrEmployeeAlreadyExists
	}

	employee.ID = res.InsertedID.(primitive.ObjectID) //nolint:forcetypeassert

	return nil
}

func (r *EmployeesRepo) Update(ctx context.Context, employee domain.UpdateEmployeeInput) error {
	updateQuery := bson.M{}

	if employee.Firstame != "" {
		updateQuery["firstname"] = employee.Firstame
	}

	if employee.Lastname != "" {
		updateQuery["lastname"] = employee.Lastname
	}

	if employee.Active != nil {
		updateQuery["active"] = *employee.Active
	}

	if employee.Status != nil {
		updateQuery["status"] = *employee.Status
	}

	if employee.Address != nil {
		updateQuery["address"] = domain.Address{
			Country:       employee.Address.Country,
			Province:      employee.Address.Province,
			City:          employee.Address.City,
			StreetAddress: employee.Address.StreetAddress,
			PostalCode:    employee.Address.PostalCode,
		}
	}

	_, err := r.db.UpdateOne(ctx,
		bson.M{"_id": employee.EmployeeID},
		bson.M{"$set": updateQuery},
	)

	return err
}

func (r *EmployeesRepo) Delete(ctx context.Context, employeeID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": employeeID})

	return err
}

func (r *EmployeesRepo) GetByCredentials(ctx context.Context, email string, password string) (domain.Employee, error) {
	var employee domain.Employee
	if err := r.db.FindOne(ctx, bson.M{
		"email":    email,
		"password": password,
	}).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Employee{}, domain.ErrEmployeeNotFound
		}

		return domain.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeesRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Employee, error) {
	var employee domain.Employee
	if err := r.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Employee{}, domain.ErrEmployeeNotFound
		}

		return domain.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeesRepo) GetById(ctx context.Context, employeeID primitive.ObjectID) (domain.Employee, error) {
	var employee domain.Employee

	if err := r.db.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Employee{}, domain.ErrEmployeeNotFound
		}

		return domain.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeesRepo) GetByWorkstation(ctx context.Context, workstationID primitive.ObjectID, query domain.GetEmployeesQuery) (domain.Employees, int64, error) {
	paginationOpts := getPaginationOpts(&query.PaginationQuery)
	paginationOpts.SetSort(bson.M{"Date": -1})
	var employees domain.Employees

	cur, err := r.db.Find(ctx, bson.M{"BiweeklyWorkstations": workstationID}, paginationOpts)
	if err != nil {
		return nil, 0, err
	}

	err = cur.All(ctx, &employees)
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountDocuments(ctx, bson.M{"BiweeklyWorkstations": workstationID})

	return employees, count, err
}

func (r *EmployeesRepo) SetSession(ctx context.Context, employeeID primitive.ObjectID, session domain.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": employeeID}, bson.M{"$set": bson.M{"session": session}})

	return err
}

func (r *EmployeesRepo) AttachSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": employeeID}, bson.M{"$addToSet": bson.M{
		"BiweeklyShifts":       shiftID,
		"BiweeklyWorkstations": workstationID,
		"BiweeklySchedules":    scheduleID,
	}})

	return err
}

func (r *EmployeesRepo) DetachSchedule(ctx context.Context, employeeID primitive.ObjectID, scheduleID primitive.ObjectID, shiftID primitive.ObjectID, workstationID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": employeeID}, bson.M{"$pull": bson.M{
		"BiweeklyShifts":       shiftID,
		"BiweeklyWorkstations": workstationID,
		"BiweeklySchedules":    scheduleID,
	}})

	return err
}
