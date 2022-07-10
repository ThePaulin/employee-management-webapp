package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Addresses []Address

type Address struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Country       string             `json:"country" bson:"country,omitempty"`
	Province      string             `json:"province" bson:"province,omitempty"`
	City          string             `json:"city" bson:"city,omitempty"`
	StreetAddress string             `json:"streetAddress" bson:"streetAddress,omitempty"`
	PostalCode    string             `json:"postalCode" bson:"postalCode,omitempty"`
}
