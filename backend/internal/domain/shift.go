package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ShiftBookingMaxDays = 14

type Shifts []Shift

type Shift struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code    int                `json:"code" bson:"code"`
	Type    string             `json:"type" bson:"type"`
	StartAt time.Time          `json:"startAt" bson:"startAt"`
	EndAt   time.Time          `json:"endAt" bson:"endAt"`
	Date    time.Time          `json:"date" bson:"date"`
	Status  string             `json:"status" bson:"status"`
}
