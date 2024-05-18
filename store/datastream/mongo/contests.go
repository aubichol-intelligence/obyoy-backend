package mongo

import (
	"context"
	"fmt"

	"obyoy-backend/model"
	storedatastream "obyoy-backend/store/datastream"
	mongoModel "obyoy-backend/store/datastream/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

// Authors handles datastream related database queries
type Authors struct {
	c *mongo.Collection
}

func (d *Authors) convertData(modelDatastream *model.Datastream) (
	mongoDatastream mongoModel.Datastream,
	err error,
) {
	err = mongoDatastream.FromModel(modelDatastream)
	return
}

// Save saves Authors from model to database
func (d *Authors) Save(modelDatastream *model.Datastream) (string, error) {
	mongoDatastream := mongoModel.Datastream{}
	var err error
	mongoDatastream, err = d.convertData(modelDatastream)
	if err != nil {
		return "", fmt.Errorf("Could not convert model datastream to mongo datastream: %w", err)
	}

	if modelDatastream.ID == "" {
		mongoDatastream.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoDatastream.ID}
	update := bson.M{"$set": mongoDatastream}
	upsert := true

	_, err = d.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoDatastream.ID.Hex(), err
}

// FindByID finds a datastream by id
func (d *Authors) FindByID(id string) (*model.Datastream, error) {
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

	datastream := mongoModel.Datastream{}
	if err := result.Decode(&datastream); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return datastream.ModelDatastream(), nil
}

// FindByID finds a datastream by id
func (d *Authors) FindNext() (*model.Datastream, error) {
	//	objectID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"is_translated": 0}

	result := d.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	datastream := mongoModel.Datastream{}

	if err := result.Decode(&datastream); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return datastream.ModelDatastream(), nil
}

// FindByDatastreamID finds a datastream by datastream id
func (d *Authors) FindByDatastreamID(id string, skip int64, limit int64) ([]*model.Datastream, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"datastream_id": objectID}

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

// CountByDatastreamID returns Authors from datastream id
func (d *Authors) CountByDatastreamID(id string) (int64, error) {
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

// FindByIDs returns all the Authors from multiple datastream ids
func (d *Authors) FindByIDs(ids ...string) ([]*model.Datastream, error) {
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
func (d *Authors) Search(text string, skip, limit int64) ([]*model.Datastream, error) {
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
func (d *Authors) FindByUser(id string, skip, limit int64) ([]*model.Datastream, error) {
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
func (d *Authors) FindByDriver(id string) (*model.Datastream, error) {
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

	datastream := mongoModel.Datastream{}
	if err := result.Decode(&datastream); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return datastream.ModelDatastream(), nil
}

// cursorToDeliveries decodes Authors one by one from the search result
func (d *Authors) cursorToDeliveries(cursor *mongo.Cursor) ([]*model.Datastream, error) {
	defer cursor.Close(context.Background())
	modelDeliveries := []*model.Datastream{}

	for cursor.Next(context.Background()) {
		datastream := mongoModel.Datastream{}
		if err := cursor.Decode(&datastream); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDeliveries = append(modelDeliveries, datastream.ModelDatastream())
	}

	return modelDeliveries, nil
}

// DeliveriesParams provides parameters for datastream specific Collection
type DeliveriesParams struct {
	dig.In
	Collection *mongo.Collection `name:"datastreams"`
}

// Store provides store for Authors
func Store(params DeliveriesParams) storedatastream.Datastreams {
	return &Authors{params.Collection}
}
