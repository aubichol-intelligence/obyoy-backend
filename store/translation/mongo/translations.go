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

// Translations handles translation related database queries
type Translations struct {
	c *mongo.Collection
}

func (d *Translations) convertData(modelTranslation *model.Translation) (
	mongoTranslation mongoModel.Translation,
	err error,
) {
	err = mongoTranslation.FromModel(modelTranslation)
	return
}

// Save saves Translations from model to database
func (d *Translations) Save(modelTranslation *model.Translation) (string, error) {
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
func (d *Translations) FindByID(id string) (*model.Translation, error) {
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
func (d *Translations) FindByTranslationID(id string, skip int64, limit int64) ([]*model.Translation, error) {
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

	return d.cursorToTranslations(cursor)
}

// CountByTranslationID returns Translations from translation id
func (d *Translations) CountByTranslationID(id string) (int64, error) {
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

// FindByIDs returns all the Translations from multiple translation ids
func (d *Translations) FindByIDs(ids ...string) ([]*model.Translation, error) {
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

	return d.cursorToTranslations(cursor)
}

// Search search for Translations given the text, skip and limit
func (d *Translations) Search(text string, skip, limit int64) ([]*model.Translation, error) {
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

	return d.cursorToTranslations(cursor)
}

// Search search for Translations given the text, skip and limit
func (d *Translations) FindByUser(id string, skip, limit int64) ([]*model.Translation, error) {
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

	return d.cursorToTranslations(cursor)
}

// Search search for Translations given the text, skip and limit
func (d *Translations) FindByDriver(id string) (*model.Translation, error) {
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

// Count returns comments from status id
func (d *Translations) Count() (int64, error) {
	//objectID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{}
	cnt, err := d.c.CountDocuments(context.Background(), filter, &options.CountOptions{})

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

// cursorToTranslations decodes Translations one by one from the search result
func (d *Translations) cursorToTranslations(cursor *mongo.Cursor) ([]*model.Translation, error) {
	defer cursor.Close(context.Background())
	modelTranslations := []*model.Translation{}

	for cursor.Next(context.Background()) {
		translation := mongoModel.Translation{}
		if err := cursor.Decode(&translation); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelTranslations = append(modelTranslations, translation.ModelTranslation())
	}

	return modelTranslations, nil
}

// TranslationsParams provides parameters for translation specific Collection
type TranslationsParams struct {
	dig.In
	Collection *mongo.Collection `name:"translations"`
}

// Store provides store for Translations
func Store(params TranslationsParams) storetranslation.Translations {
	return &Translations{params.Collection}
}
