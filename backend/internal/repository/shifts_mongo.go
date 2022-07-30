package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShiftsRepo struct {
	db *mongo.Collection
}

func NewShiftsRepo(db *mongo.Database) *ShiftsRepo {
	return &ShiftsRepo{
		db: db.Collection(shiftsCollection),
	}
}

func (r *ShiftsRepo) Create(ctx context.Context, shift domain.Shift) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, shift)
	shift.ID = res.InsertedID.(primitive.ObjectID)

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *ShiftsRepo) Update(ctx context.Context, shiftID primitive.ObjectID, input UpdateShiftInput) error {
	updateQuery := bson.M{}

	if input.Type != nil {
		updateQuery["type"] = *input.Type
	}

	if input.StartAt != nil {
		updateQuery["startAt"] = *input.StartAt
	}

	if input.EndAt != nil {
		updateQuery["endAt"] = *input.EndAt
	}

	if input.Status != nil {
		updateQuery["status"] = *input.Status
	}

	_, err := r.db.UpdateOne(ctx,
		bson.M{"_id": shiftID},
		bson.M{"$set": updateQuery},
	)

	return err
}

func (r *ShiftsRepo) Delete(ctx context.Context, shiftID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": shiftID})

	return err
}

func (r *ShiftsRepo) GetById(ctx context.Context, shiftID primitive.ObjectID) (domain.Shift, error) {
	var shift domain.Shift

	if err := r.db.FindOne(ctx, bson.M{"_id": shiftID}).Decode(&shift); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Shift{}, domain.ErrShiftNotFound
		}

		return domain.Shift{}, err
	}

	return shift, nil
}

func (r *ShiftsRepo) GetByStatus(ctx context.Context, shiftID primitive.ObjectID, status string, query domain.GetShiftsQuery) (domain.Shifts, int64, error) {
	paginationOpts := getPaginationOpts(&query.PaginationQuery)
	paginationOpts.SetSort(bson.M{"date": -1})
	var shifts domain.Shifts

	cur, err := r.db.Find(ctx, bson.M{"status": status}, paginationOpts)
	if err != nil {
		return nil, 0, err
	}

	err = cur.All(ctx, &shifts)
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountDocuments(ctx, bson.M{"status": status})

	return shifts, count, err
}

func (r *ShiftsRepo) SetStatus(ctx context.Context, shiftID primitive.ObjectID, status string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": shiftID}, bson.M{"$set": bson.M{"status": status}})

	return err
}
