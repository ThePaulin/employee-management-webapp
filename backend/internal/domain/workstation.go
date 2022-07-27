package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const WorkstationBookingMaxDays = 14

type Workstations []Workstation

type Workstation struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Code           int                  `json:"code" bson:"code,omitempty"`
	Name           string               `json:"name" bson:"name"`
	Description    string               `json:"description" bson:"description,omitempty"`
	Status         string               `json:"status" bson:"status,omitempty"`
	RegisteredAt   time.Time            `json:"registeredAt" bson:"registeredAtm,omitempty"`
	Managers       []primitive.ObjectID `json:"managers" bson:"managers,omitempty"`
	BiweeklyShifts []primitive.ObjectID `json:"biweeklyShifts" bson:"biweeklyShifts,omitempty"`
	ShiftsHistory  []primitive.ObjectID `json:"shiftsHistory" bson:"shiftsHistory,omitempty"`
}
