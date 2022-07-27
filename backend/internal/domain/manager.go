package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Managers []Manager

type Manager struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Permission   string               `json:"permission" bson:"permission"`
	Firstame     string               `json:"firstname" bson:"firstname,omitempty"`
	Lastname     string               `json:"lastname" bson:"lastname,omitempty"`
	Email        string               `json:"email" bson:"email"`
	Password     string               `json:"password" bson:"password"`
	Session      Session              `json:"session" bson:"session,omitempty"`
	Workstations []primitive.ObjectID `json:"workstations" bson:"workstations"`
}
