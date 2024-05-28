package model

import (
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Datastream holds db data type for datastreams
type Datastream struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UploaderID      primitive.ObjectID `bson:"uploader_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	SourceSentence  string             `bson:"source_sentence,omitempty"`
	SourceLanguage  string             `bson:"source_language,omitempty"`
	LineNumber      int32              `bson:"line_number,omitempty"`
	DatasetID       primitive.ObjectID `bson:"dataset_id,omitempty"`
	IsTranslated    int32              `bson:"is_translated,omitempty"`
	TimesTranslated int32              `bson:"times_translated,omitempty"`
	TimesReviewed   int32              `bson:"times_reviewed,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	IsDeleted       bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for datastreams
func (d *Datastream) FromModel(modelDatastream *model.Datastream) error {
	d.CreatedAt = modelDatastream.CreatedAt
	d.UpdatedAt = modelDatastream.UpdatedAt
	d.SourceSentence = modelDatastream.SourceSentence
	d.SourceLanguage = modelDatastream.SourceLanguage
	d.LineNumber = modelDatastream.LineNumber
	d.TimesTranslated = modelDatastream.TimesTranslated
	d.TimesReviewed = modelDatastream.TimesReviewed
	d.IsTranslated = modelDatastream.IsTranslated
	d.IsDeleted = modelDatastream.IsDeleted
	var err error

	if modelDatastream.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelDatastream.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	if modelDatastream.DatasetID != "" {
		d.DatasetID, err = primitive.ObjectIDFromHex(modelDatastream.DatasetID)
	}

	if err != nil {
		return err
	}

	if modelDatastream.UploaderID != "" {
		d.UploaderID, err = primitive.ObjectIDFromHex(modelDatastream.UploaderID)
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelDatastream converts bson to model
func (d *Datastream) ModelDatastream() *model.Datastream {
	Datastream := model.Datastream{}

	Datastream.ID = d.ID.Hex()
	Datastream.DatasetID = d.DatasetID.Hex()
	Datastream.UploaderID = d.UploaderID.Hex()
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
