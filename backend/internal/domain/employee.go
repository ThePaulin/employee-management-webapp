package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employees []Employee

type Employee struct {
	ID                primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Active            bool                 `json:"active" bson:"active"`
	Firstame          string               `json:"firstname" bson:"firstname,omitempty"`
	Lastname          string               `json:"lastname" bson:"lastname,omitempty"`
	Username          string               `json:"username" bson:"username,omitempty"`
	Email             string               `json:"email" bson:"email,omitempty"`
	Password          string               `json:"password" bson:"password,omitempty"`
	PhoneNumber       string               `json:"phoneNumber" bson:"phoneNumber,omitempty"`
	Address           Address              `json:"addresses" bson:"addresses,omitempty"`
	Status            string               `json:"status" bson:"status,omitempty"`
	Date              time.Time            `json:"date" bson:"date,omitempty"`
	BiweeklySchedules []primitive.ObjectID `json:"biweeklySchedules" bson:"biweeklySchedules"`
	SchedulesHistory  []primitive.ObjectID `json:"schedulesHistory" bson:"schedulesHistory"`
}

type Address struct {
	Country       string `json:"country" bson:"country,omitempty"`
	Province      string `json:"province" bson:"province,omitempty"`
	City          string `json:"city" bson:"city,omitempty"`
	StreetAddress string `json:"streetAddress" bson:"streetAddress,omitempty"`
	PostalCode    string `json:"postalCode" bson:"postalCode,omitempty"`
}

func (emp Employee) IsBiweeklySchedules() bool {
	return len(emp.BiweeklySchedules) < ScheduleMaxDays
}

func (emp Employee) IsScheduleAssigned(s Schedule) bool {
	for _, id := range emp.BiweeklySchedules {
		if s.ID == id {
			return true
		}
	}
	return false
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
	EmployeeID        primitive.ObjectID   `json:"_"`
	Firstame          string               `json:"firstname"`
	Lastname          string               `json:"lastname"`
	Address           Address              `json:"address"`
	Active            *bool                `json:"active"`
	Status            *string              `json:"status"`
	BiweeklySchedules []primitive.ObjectID `json:"biweeklySchedules"`
}

type CreateEmployeeInput struct {
	Firstame string  `json:"firstname" binding:"required,min=2"`
	Lastname string  `json:"lastname" binding:"required,min=2"`
	Username string  `json:"username" binding:"required,min=6"`
	Password string  `json:"password" binding:"required,min=6"`
	Email    string  `json:"email" binding:"required,email"`
	Active   bool    `json:"active" binding:"required"`
	Address  Address `json:"address"`
}
