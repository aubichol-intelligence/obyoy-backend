package mongo

import (
	"context"
	"fmt"

	"ardent-backend/model"
	storecontest "ardent-backend/store/contest"
	mongoModel "ardent-backend/store/contest/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

// Authors handles contest related database queries
type Authors struct {
	c *mongo.Collection
}

func (d *Authors) convertData(modelContest *model.Contest) (
	mongoContest mongoModel.Contest,
	err error,
) {
	err = mongoContest.FromModel(modelContest)
	return
}

// Save saves Authors from model to database
func (d *Authors) Save(modelContest *model.Contest) (string, error) {
	mongoContest := mongoModel.Contest{}
	var err error
	mongoContest, err = d.convertData(modelContest)
	if err != nil {
		return "", fmt.Errorf("Could not convert model contest to mongo contest: %w", err)
	}

	if modelContest.ID == "" {
		mongoContest.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoContest.ID}
	update := bson.M{"$set": mongoContest}
	upsert := true

	_, err = d.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoContest.ID.Hex(), err
}

// FindByID finds a contest by id
func (d *Authors) FindByID(id string) (*model.Contest, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := d.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	contest := mongoModel.Contest{}
	if err := result.Decode(&contest); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return contest.ModelContest(), nil
}

// FindByContestID finds a contest by contest id
func (d *Authors) FindByContestID(id string, skip int64, limit int64) ([]*model.Contest, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"contest_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := d.c.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	return d.cursorToDeliveries(cursor)
}

// CountByContestID returns Authors from contest id
func (d *Authors) CountByContestID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}
	cnt, err := d.c.CountDocuments(context.Background(), filter, &options.CountOptions{})

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

// FindByIDs returns all the Authors from multiple contest ids
func (d *Authors) FindByIDs(ids ...string) ([]*model.Contest, error) {
	objectIDs := []primitive.ObjectID{}
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("Invalid id %s : %w", id, err)
		}

		objectIDs = append(objectIDs, objectID)
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": objectIDs,
		},
	}

	cursor, err := d.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return d.cursorToDeliveries(cursor)
}

// Search search for Authors given the text, skip and limit
func (d *Authors) Search(text string, skip, limit int64) ([]*model.Contest, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := d.c.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		},
	)
	if err != nil {
		return nil, err
	}

	return d.cursorToDeliveries(cursor)
}

// Search search for Authors given the text, skip and limit
func (d *Authors) FindByUser(id string, skip, limit int64) ([]*model.Contest, error) {
	filter := bson.M{"_id": id}
	cursor, err := d.c.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		},
	)
	if err != nil {
		return nil, err
	}

	return d.cursorToDeliveries(cursor)
}

// Search search for Authors given the text, skip and limit
func (d *Authors) FindByDriver(id string) (*model.Contest, error) {
	driverID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"driver_id": driverID, "is_active": true, "state": "pending"}

	result := d.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)

	if err := result.Err(); err != nil {
		return nil, err
	}

	contest := mongoModel.Contest{}
	if err := result.Decode(&contest); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return contest.ModelContest(), nil
}

// cursorToDeliveries decodes Authors one by one from the search result
func (d *Authors) cursorToDeliveries(cursor *mongo.Cursor) ([]*model.Contest, error) {
	defer cursor.Close(context.Background())
	modelDeliveries := []*model.Contest{}

	for cursor.Next(context.Background()) {
		contest := mongoModel.Contest{}
		if err := cursor.Decode(&contest); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDeliveries = append(modelDeliveries, contest.ModelContest())
	}

	return modelDeliveries, nil
}

// DeliveriesParams provides parameters for contest specific Collection
type DeliveriesParams struct {
	dig.In
	Collection *mongo.Collection `name:"contests"`
}

// Store provides store for Authors
func Store(params DeliveriesParams) storecontest.Contests {
	return &Authors{params.Collection}
}
