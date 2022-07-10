package repository

import "go.mongodb.org/mongo-driver/mongo"

type EmployeesRepo struct {
	db *mongo.Collection
}

func NewEmployeesRepo(db *mongo.Database) *EmployeesRepo {
	return &EmployeesRepo{
		db: db.Collection(employeesCollection),
	}
}
