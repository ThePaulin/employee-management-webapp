package domain

import "time"

type Session struct {
	RefreshedToken string    `json:"refreshToken" bson:"refreshToken"`
	ExpDate        time.Time `json:"expDate" bson:"expDate"`
}
