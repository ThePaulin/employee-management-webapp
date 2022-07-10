package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Managers []Manager

type Manager struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Permission string             `json:"permission" bson:"permission"`
	Person     Person             `json:"person" bson:"person,omitempty"`
}
