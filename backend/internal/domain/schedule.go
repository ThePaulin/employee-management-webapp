package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ScheduleMaxDays = 14

type Schedules []Schedule

type Schedule struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EmployeeID    primitive.ObjectID `json:"employeeId" bson:"employee,omitempty"`
	ShiftID       primitive.ObjectID `json:"shiftId" bson:"shift,omitempty"`
	WorkstationID primitive.ObjectID `json:"workstationId" bson:"shift,omitempty"`
	Date          time.Time          `json:"date" bson:"date,omitempty"`
}

type ScheduleInfo struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Name string             `json:"name" bson:"name"`
}

type ScheduleShiftInfo struct {
	ID   primitive.ObjectID `json:"id" bson:"id"`
	Code string             `json:"code" bson:"code"`
}
