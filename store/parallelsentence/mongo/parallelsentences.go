package mongo

import (
	"context"
	"fmt"

	"obyoy-backend/model"
	storeparallelsentence "obyoy-backend/store/parallelsentence"
	mongoModel "obyoy-backend/store/parallelsentence/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

// Authors handles parallelsentence related database queries
type Authors struct {
	c *mongo.Collection
}

func (d *Authors) convertData(modelParallelsentence *model.Parallelsentence) (
	mongoParallelsentence mongoModel.Parallelsentence,
	err error,
) {
	err = mongoParallelsentence.FromModel(modelParallelsentence)
	return
}

// Save saves Authors from model to database
func (d *Authors) Save(modelParallelsentence *model.Parallelsentence) (string, error) {
	mongoParallelsentence := mongoModel.Parallelsentence{}
	var err error
	mongoParallelsentence, err = d.convertData(modelParallelsentence)
	if err != nil {
		return "", fmt.Errorf("Could not convert model parallelsentence to mongo parallelsentence: %w", err)
	}

	filter := bson.M{"_id": mongoParallelsentence.ID}
	update := bson.M{"$set": mongoParallelsentence}
	upsert := true

	_, err = d.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoParallelsentence.ID.Hex(), err
}

// FindByID finds a parallelsentence by id
func (d *Authors) FindByID(id string) (*model.Parallelsentence, error) {
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

	parallelsentence := mongoModel.Parallelsentence{}
	if err := result.Decode(&parallelsentence); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return parallelsentence.ModelParallelsentence(), nil
}

// FindByParallelsentenceID finds a parallelsentence by parallelsentence id
func (d *Authors) FindByParallelsentenceID(id string, skip int64, limit int64) ([]*model.Parallelsentence, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"parallelsentence_id": objectID}

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

// CountByParallelsentenceID returns Authors from parallelsentence id
func (d *Authors) CountByParallelsentenceID(id string) (int64, error) {
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

// FindByIDs returns all the Authors from multiple parallelsentence ids
func (d *Authors) FindByIDs(ids ...string) ([]*model.Parallelsentence, error) {
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
func (d *Authors) Search(text string, skip, limit int64) ([]*model.Parallelsentence, error) {
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
func (d *Authors) FindByUser(id string, skip, limit int64) ([]*model.Parallelsentence, error) {
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
func (d *Authors) FindByDriver(id string) (*model.Parallelsentence, error) {
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

	parallelsentence := mongoModel.Parallelsentence{}
	if err := result.Decode(&parallelsentence); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return parallelsentence.ModelParallelsentence(), nil
}

// cursorToDeliveries decodes Authors one by one from the search result
func (d *Authors) cursorToDeliveries(cursor *mongo.Cursor) ([]*model.Parallelsentence, error) {
	defer cursor.Close(context.Background())
	modelDeliveries := []*model.Parallelsentence{}

	for cursor.Next(context.Background()) {
		parallelsentence := mongoModel.Parallelsentence{}
		if err := cursor.Decode(&parallelsentence); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDeliveries = append(modelDeliveries, parallelsentence.ModelParallelsentence())
	}

	return modelDeliveries, nil
}

// DeliveriesParams provides parameters for parallelsentence specific Collection
type DeliveriesParams struct {
	dig.In
	Collection *mongo.Collection `name:"parallelsentences"`
}

// Store provides store for Authors
func Store(params DeliveriesParams) storeparallelsentence.Parallelsentences {
	return &Authors{params.Collection}
}