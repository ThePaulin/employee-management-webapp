package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Persons []Person

type Person struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstame    string             `json:"firstname" bson:"firstname,omitempty"`
	Lastname    string             `json:"lastname" bson:"lastname,omitempty"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber,omitempty"`
	Addresses   Addresses          `json:"addresses" bson:"addresses,omitempty"`
}
