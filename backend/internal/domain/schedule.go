package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ScheduleMaxDays = 14

type Schedules []Schedule

type Schedule struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Employee Employee           `json:"employee" bson:"employee,omitempty"`
	Shift    Shift              `json:"shift" bson:"shift,omitempty"`
	Date     time.Time          `json:"date" bson:"date,omitempty"`
}
