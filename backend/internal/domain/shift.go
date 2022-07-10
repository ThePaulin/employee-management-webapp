package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shifts []Shift

type Shift struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code    int                `json:"code" bson:"code"`
	Type    string             `json:"permission" bson:"permissiony"`
	StartAt time.Time          `json:"startAt" bson:"startAt"`
	EndAt   time.Time          `json:"endAt" bson:"endAt"`
	Status  string             `json:"status" bson:"status"`
}
