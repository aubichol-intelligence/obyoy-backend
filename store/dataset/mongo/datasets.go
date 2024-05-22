package mongo

import (
	"context"
	"fmt"

	"obyoy-backend/model"
	storedataset "obyoy-backend/store/dataset"
	mongoModel "obyoy-backend/store/dataset/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

// Authors handles dataset related database queries
type Authors struct {
	c *mongo.Collection
}

func (d *Authors) convertData(modelDataset *model.Dataset) (
	mongoDataset mongoModel.Dataset,
	err error,
) {
	err = mongoDataset.FromModel(modelDataset)
	return
}

// Save saves Authors from model to database
func (d *Authors) Save(modelDataset *model.Dataset) (string, error) {
	mongoDataset := mongoModel.Dataset{}
	var err error
	mongoDataset, err = d.convertData(modelDataset)
	if err != nil {
		return "", fmt.Errorf("Could not convert model dataset to mongo dataset: %w", err)
	}

	if modelDataset.ID == "" {
		mongoDataset.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoDataset.ID}
	update := bson.M{"$set": mongoDataset}
	upsert := true

	_, err = d.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoDataset.ID.Hex(), err
}

// FindByID finds a dataset by id
func (d *Authors) FindByID(id string) (*model.Dataset, error) {
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

	dataset := mongoModel.Dataset{}
	if err := result.Decode(&dataset); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return dataset.ModelDataset(), nil
}

// FindByDatasetID finds a dataset by dataset id
func (d *Authors) FindByDatasetID(id string, skip int64, limit int64) ([]*model.Dataset, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"dataset_id": objectID}

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

// CountByDatasetID returns Authors from dataset id
func (d *Authors) CountByDatasetID(id string) (int64, error) {
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

// FindByIDs returns all the Authors from multiple dataset ids
func (d *Authors) FindByIDs(ids ...string) ([]*model.Dataset, error) {
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
func (d *Authors) Search(text string, skip, limit int64) ([]*model.Dataset, error) {
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
func (d *Authors) FindByUser(id string, skip, limit int64) ([]*model.Dataset, error) {
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
func (d *Authors) FindByDriver(id string) (*model.Dataset, error) {
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

	dataset := mongoModel.Dataset{}
	if err := result.Decode(&dataset); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return dataset.ModelDataset(), nil
}

// cursorToDeliveries decodes Authors one by one from the search result
func (d *Authors) cursorToDeliveries(cursor *mongo.Cursor) ([]*model.Dataset, error) {
	defer cursor.Close(context.Background())
	modelDeliveries := []*model.Dataset{}

	for cursor.Next(context.Background()) {
		dataset := mongoModel.Dataset{}
		if err := cursor.Decode(&dataset); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDeliveries = append(modelDeliveries, dataset.ModelDataset())
	}

	return modelDeliveries, nil
}

// DeliveriesParams provides parameters for dataset specific Collection
type DeliveriesParams struct {
	dig.In
	Collection *mongo.Collection `name:"datasets"`
}

// Store provides store for Authors
func Store(params DeliveriesParams) storedataset.Datasets {
	return &Authors{params.Collection}
}