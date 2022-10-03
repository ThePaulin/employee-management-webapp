package service

import (
	"context"
	"employee-management-webapp/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManagerLoginInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Managers interface {
	Login(ctx context.Context, input ManagerLoginInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, managerID primitive.ObjectID, hash string) error
	CreateWorkstation(ctx context.Context, managerID primitive.ObjectID, workstationName string, workstationCode string) (domain.Workstation, error)
}

type ConnectSendPulseInput struct {
	WorkstationID primitive.ObjectID
	ID            string
	Secret        string
	ListID        string
}

type Workstations interface {
	Create(ctx context.Context, name string) (primitive.ObjectID, error)
	GetByDomain(ctx context.Context, domainName string) (domain.Workstation, error)
	GetById(ctx context.Context, id primitive.ObjectID) (domain.Workstation, error)
	ConnectSendPulse(ctx context.Context, input ConnectSendPulseInput) error
}
