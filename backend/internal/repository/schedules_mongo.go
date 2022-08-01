package repository

import (
	"context"
	"employee-management-webapp/internal/domain"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchedulesRepo struct {
	db *mongo.Collection
}

func NewSchedulesRepo(db *mongo.Database) *SchedulesRepo {
	return &SchedulesRepo{
		db: db.Collection(schedulesCollection),
	}
}

func (r *SchedulesRepo) Create(ctx context.Context, schedule domain.Schedule) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, schedule)
	schedule.ID = res.InsertedID.(primitive.ObjectID)

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *SchedulesRepo) Update(ctx context.Context, scheduleID primitive.ObjectID, employeeID primitive.ObjectID) error {
	var err error
	if scheduleID != primitive.NilObjectID && employeeID != primitive.NilObjectID {
		_, err = r.db.UpdateOne(ctx, bson.M{"_id": scheduleID}, bson.M{"$set": bson.M{"employeeId": employeeID}})
	}

	return err
}

func (r *SchedulesRepo) Delete(ctx context.Context, scheduleID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": scheduleID})

	return err
}

func (r *SchedulesRepo) GetByEmployee(ctx context.Context, employeeID primitive.ObjectID, query domain.GetSchedulesQuery) (domain.Schedules, int64, error) {
	paginationOpts := getPaginationOpts(&query.PaginationQuery)
	paginationOpts.SetSort(bson.M{"date": -1})
	var schedules domain.Schedules

	cur, err := r.db.Find(ctx, bson.M{"employeeId": employeeID}, paginationOpts)
	if err != nil {
		return nil, 0, err
	}

	err = cur.All(ctx, &schedules)
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountDocuments(ctx, bson.M{"employeeId": employeeID})

	return schedules, count, err
}

func (r *SchedulesRepo) GetById(ctx context.Context, scheduleID primitive.ObjectID) (domain.Schedule, error) {
	var schedule domain.Schedule

	if err := r.db.FindOne(ctx, bson.M{"_id": scheduleID}).Decode(&schedule); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Schedule{}, domain.ErrScheduleNotFound
		}

		return domain.Schedule{}, err
	}

	return schedule, nil
}

func (r *SchedulesRepo) GetByShift(ctx context.Context, shiftID primitive.ObjectID) (domain.Schedule, error) {
	var schedule domain.Schedule

	if err := r.db.FindOne(ctx, bson.M{"shiftId": shiftID}).Decode(&schedule); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Schedule{}, domain.ErrScheduleNotFound
		}

		return domain.Schedule{}, err
	}

	return schedule, nil
}

func (r *SchedulesRepo) GetByIds(ctx context.Context, scheduleIDs []primitive.ObjectID) (domain.Schedules, error) {
	var schedules domain.Schedules

	cur, err := r.db.Find(ctx, bson.M{"_id": bson.M{"$in": scheduleIDs}})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &schedules)

	return schedules, err
}
