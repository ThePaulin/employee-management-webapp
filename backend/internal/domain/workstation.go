package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workstations []Workstation

type Workstation struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code   int                `json:"code" bson:"code"`
	Name   string             `json:"name" bson:"name"`
	Status string             `json:"status" bson:"status"`
}
