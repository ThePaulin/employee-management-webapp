package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkstationsRepo struct {
	db *mongo.Collection
}

func NewWorkstationsRepo(db *mongo.Database) *WorkstationsRepo {
	return &WorkstationsRepo{
		db: db.Collection(workstationsCollection),
	}
}

func (r *WorkstationsRepo) Create(ctx context.Context, code int, name string) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, domain.Workstation{
		Name:         name,
		Code:         code,
		RegisteredAt: time.Now(),
	})

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *WorkstationsRepo) GetById(ctx context.Context, workstationID primitive.ObjectID) (domain.Workstation, error) {
	var workstation domain.Workstation

	if err := r.db.FindOne(ctx, bson.M{"_id": workstationID}).Decode(&workstation); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Workstation{}, domain.ErrEmployeeNotFound
		}

		return domain.Workstation{}, err
	}

	return workstation, nil
}

func (r *WorkstationsRepo) GetByManager(ctx context.Context, managerID primitive.ObjectID) (domain.Workstations, error) {
	var workstations domain.Workstations

	cur, err := r.db.Find(ctx, bson.M{"managers": managerID})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &workstations)
	if err != nil {
		return nil, err
	}

	return workstations, err
}

func (r *WorkstationsRepo) GetByShift(ctx context.Context, shiftID primitive.ObjectID) (domain.Workstation, error) {
	var workstation domain.Workstation

	cur, err := r.db.Find(ctx, bson.M{"biweeklyShifts": shiftID})
	if err != nil {
		return domain.Workstation{}, err
	}

	err = cur.All(ctx, &workstation)
	if err != nil {
		return domain.Workstation{}, err
	}

	return workstation, err
}

func (r *WorkstationsRepo) Update(ctx context.Context, workstationID primitive.ObjectID, workstation UpdateWorkstationInput) error {
	updateQuery := bson.M{}

	if workstation.Name != nil {
		updateQuery["name"] = workstation.Name
	}

	if workstation.Description != nil {
		updateQuery["description"] = workstation.Description
	}

	if workstation.Status != nil {
		updateQuery["status"] = workstation.Status
	}

	_, err := r.db.UpdateOne(ctx,
		bson.M{"_id": workstationID},
		bson.M{"$set": updateQuery},
	)

	return err
}
