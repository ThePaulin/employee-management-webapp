package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employees []Employee

type Employee struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Active    bool               `json:"active" bson:"active"`
	Person    Person             `json:"person" bson:"person"`
	Status    bool               `json:"status" bson:"status"`
	Scheduled Schedules          `json:"schedules" bson:"schedules"`
	Finished  Schedules          `json:"finished" bson:"finished"`
}

func (emp Employee) IsFullyScheduled() bool {
	return len(emp.Scheduled) < ScheduleMaxDays
}

type EmployeeSchedules struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id"`
	Ongoing  []primitive.ObjectID `json:"ongoing" bson:"ongoing"`
	Finished []primitive.ObjectID `json:"finished" bson:"finished"`
}

type EmployeeInfoShort struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Firstame    string             `json:"firstname" bson:"firstname"`
	Lastname    string             `json:"lastname" bson:"lastname"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Address     Address            `json:"address" bson:"address"`
	Active      bool               `json:"active" bson:"active"`
}

type UpdateEmployeeInput struct {
	Firstame    string               `json:"firstname"`
	Lastname    string               `json:"lastname"`
	Email       string               `json:"email"`
	PhoneNumber string               `json:"phoneNumber"`
	Address     Address              `json:"address"`
	Active      bool                 `json:"active"`
	EmployeeID  primitive.ObjectID   `json:"_"`
	Scheduled   []primitive.ObjectID `json:"scheduled"`
}

type CreateEmployeeInput struct {
	Firstame  string               `json:"firstname" binding:"required,min=2"`
	Lastname  string               `json:"lastname" binding:"required,min=2"`
	Username  string               `json:"username" binding:"required,min=6"`
	Password  string               `json:"password" binding:"required,min=6"`
	Email     string               `json:"email" binding:"required,email"`
	Active    bool                 `json:"active" binding:"required"`
	Addresses []primitive.ObjectID `json:"addresses"`
}
