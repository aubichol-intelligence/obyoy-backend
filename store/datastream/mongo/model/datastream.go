package model

import (
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Datastream holds db data type for datastreams
type Datastream struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	SourceSentence  string             `bson:"source_sentence,omitempty"`
	SourceLanguage  string             `bson:"source_language,omitempty"`
	LineNumber      int32              `bson:"line_number,omitempty"`
	DatasetID       primitive.ObjectID `bson:"dataset_id,omitempty"`
	IsTranslated    int32              `bson:"is_translated"`
	TimesTranslated int32              `bson:"times_translated,omitempty"`
	TimesReviewed   int32              `bson:"times_reviewed,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	IsDeleted       bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Datastream) FromModel(modelDelivery *model.Datastream) error {
	d.CreatedAt = modelDelivery.CreatedAt
	d.UpdatedAt = modelDelivery.UpdatedAt
	d.SourceSentence = modelDelivery.SourceSentence
	d.SourceLanguage = modelDelivery.SourceLanguage
	d.LineNumber = modelDelivery.LineNumber
	d.TimesTranslated = modelDelivery.TimesTranslated
	d.TimesReviewed = modelDelivery.TimesReviewed
	d.IsTranslated = modelDelivery.IsTranslated

	var err error

	if modelDelivery.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelDelivery.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	if modelDelivery.DatasetID != "" {
		d.DatasetID, err = primitive.ObjectIDFromHex(modelDelivery.DatasetID)
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelDelivery converts bson to model
func (d *Datastream) ModelDatastream() *model.Datastream {
	Datastream := model.Datastream{}
	Datastream.ID = d.ID.Hex()
	Datastream.DatasetID = d.DatasetID.Hex()
	Datastream.CreatedAt = d.CreatedAt
	Datastream.UpdatedAt = d.UpdatedAt
	Datastream.SourceSentence = d.SourceSentence
	Datastream.SourceLanguage = d.SourceLanguage
	Datastream.LineNumber = d.LineNumber
	Datastream.IsTranslated = d.IsTranslated
	Datastream.TimesTranslated = d.TimesTranslated
	Datastream.TimesReviewed = d.TimesReviewed

	return &Datastream
}
