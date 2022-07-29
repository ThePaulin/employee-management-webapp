package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"employee-management-webapp/pkg/database/mongodb"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManagersRepo struct {
	db *mongo.Collection
}

func NewManagersRepo(db *mongo.Database) *ManagersRepo {
	return &ManagersRepo{
		db: db.Collection(managersCollection),
	}
}

func (r *ManagersRepo) Create(ctx context.Context, manager *domain.Manager) error {
	res, err := r.db.InsertOne(ctx, manager)
	if mongodb.IsDuplicate(err) {
		return domain.ErrManagerAlreadyExists
	}

	manager.ID = res.InsertedID.(primitive.ObjectID) //nolint:forcetypeassert

	return nil
}

func (r *ManagersRepo) GetByCredentials(ctx context.Context, email string, password string) (domain.Manager, error) {
	var manager domain.Manager
	if err := r.db.FindOne(ctx, bson.M{
		"email":    email,
		"password": password,
	}).Decode(&manager); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Manager{}, domain.ErrManagerNotFound
		}

		return domain.Manager{}, err
	}

	return manager, nil
}

func (r *ManagersRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Manager, error) {
	var manager domain.Manager
	if err := r.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&manager); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Manager{}, domain.ErrManagerNotFound
		}

		return domain.Manager{}, err
	}

	return manager, nil
}

func (r *ManagersRepo) SetSession(ctx context.Context, managerID primitive.ObjectID, session domain.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": managerID}, bson.M{"$set": bson.M{"session": session}})

	return err
}

func (r *ManagersRepo) AttachWorkstation(ctx context.Context, managerID primitive.ObjectID, workstationID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": managerID}, bson.M{"$addToSet": bson.M{
		"workstations": workstationID,
	}})

	return err
}
