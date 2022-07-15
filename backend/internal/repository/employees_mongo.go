package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"employee-management-webapp/pkg/database/mongodb"
	"reflect"

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
		updateQuery["active"] = employee.Active
	}

	// if employee.Status != "" {
	// 	updateQuery["status"] = employee.Active
	// }

	if (reflect.TypeOf(employee.Address) == reflect.TypeOf(domain.Address{}) && employee.Address != domain.Address{}) {
		updateQuery["address"] = domain.Address{
			ID:            employee.Address.ID,
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
