package mongo

import (
	"context"
	"fmt"

	"obyoy-backend/model"
	storetranslation "obyoy-backend/store/translation"
	mongoModel "obyoy-backend/store/translation/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

// Authors handles translation related database queries
type Authors struct {
	c *mongo.Collection
}

func (d *Authors) convertData(modelTranslation *model.Translation) (
	mongoTranslation mongoModel.Translation,
	err error,
) {
	err = mongoTranslation.FromModel(modelTranslation)
	return
}

// Save saves Authors from model to database
func (d *Authors) Save(modelTranslation *model.Translation) (string, error) {
	mongoTranslation := mongoModel.Translation{}
	var err error
	mongoTranslation, err = d.convertData(modelTranslation)
	if err != nil {
		return "", fmt.Errorf("Could not convert model translation to mongo translation: %w", err)
	}

	if modelTranslation.ID == "" {
		mongoTranslation.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoTranslation.ID}
	update := bson.M{"$set": mongoTranslation}
	upsert := true

	_, err = d.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoTranslation.ID.Hex(), err
}

// FindByID finds a translation by id
func (d *Authors) FindByID(id string) (*model.Translation, error) {
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

	translation := mongoModel.Translation{}
	if err := result.Decode(&translation); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return translation.ModelTranslation(), nil
}

// FindByTranslationID finds a translation by translation id
func (d *Authors) FindByTranslationID(id string, skip int64, limit int64) ([]*model.Translation, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"translation_id": objectID}

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

// CountByTranslationID returns Authors from translation id
func (d *Authors) CountByTranslationID(id string) (int64, error) {
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

// FindByIDs returns all the Authors from multiple translation ids
func (d *Authors) FindByIDs(ids ...string) ([]*model.Translation, error) {
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
func (d *Authors) Search(text string, skip, limit int64) ([]*model.Translation, error) {
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
func (d *Authors) FindByUser(id string, skip, limit int64) ([]*model.Translation, error) {
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
func (d *Authors) FindByDriver(id string) (*model.Translation, error) {
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

	translation := mongoModel.Translation{}
	if err := result.Decode(&translation); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return translation.ModelTranslation(), nil
}

// cursorToDeliveries decodes Authors one by one from the search result
func (d *Authors) cursorToDeliveries(cursor *mongo.Cursor) ([]*model.Translation, error) {
	defer cursor.Close(context.Background())
	modelDeliveries := []*model.Translation{}

	for cursor.Next(context.Background()) {
		translation := mongoModel.Translation{}
		if err := cursor.Decode(&translation); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDeliveries = append(modelDeliveries, translation.ModelTranslation())
	}

	return modelDeliveries, nil
}

// DeliveriesParams provides parameters for translation specific Collection
type DeliveriesParams struct {
	dig.In
	Collection *mongo.Collection `name:"translations"`
}

// Store provides store for Authors
func Store(params DeliveriesParams) storetranslation.Translations {
	return &Authors{params.Collection}
}
