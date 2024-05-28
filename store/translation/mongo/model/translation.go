package model

import (
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Translation holds db data type for translations
type Translation struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	DatasetID           primitive.ObjectID `bson:"dataset_id,omitempty"`
	DatastreamID        primitive.ObjectID `bson:"datastream_id,omitempty"`
	TranslatorID        primitive.ObjectID `bson:"translator_id,omitempty"`
	ReviewerID          primitive.ObjectID `bson:"reviewer_id,omitempty"`
	CreatedAt           time.Time          `bson:"created_at,omitempty"`
	UpdatedAt           time.Time          `bson:"updated_at,omitempty"`
	SourceSentence      string             `bson:"source_sentence,omitempty"`
	SourceLanguage      string             `bson:"source_language,omitempty"`
	DestinationSentence string             `bson:"destination_sentence,omitempty"`
	DestinationLanguage string             `bson:"destination_language,omitempty"`
	Name                string             `bson:"name,omitempty"`
	LineNumber          int                `bson:"line_number,omitempty"`
	IsDeleted           bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for translations
func (d *Translation) FromModel(modelTranslation *model.Translation) error {
	d.CreatedAt = modelTranslation.CreatedAt
	d.UpdatedAt = modelTranslation.UpdatedAt
	d.SourceLanguage = modelTranslation.SourceLanguage
	d.SourceSentence = modelTranslation.SourceSentence
	d.DestinationSentence = modelTranslation.DestinationSentence
	d.DestinationLanguage = modelTranslation.DestinationLanguage
	d.Name = modelTranslation.Name
	d.LineNumber = modelTranslation.LineNumber
	d.IsDeleted = modelTranslation.IsDeleted

	var err error

	if modelTranslation.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelTranslation.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	if modelTranslation.DatasetID != "" {
		d.DatasetID, err = primitive.ObjectIDFromHex(modelTranslation.DatasetID)
	}

	if err != nil {
		return err
	}

	if modelTranslation.DatastreamID != "" {
		d.DatastreamID, err = primitive.ObjectIDFromHex(modelTranslation.DatastreamID)
	}

	if err != nil {
		return err
	}

	if modelTranslation.TranslatorID != "" {
		d.TranslatorID, err = primitive.ObjectIDFromHex(modelTranslation.TranslatorID)
	}

	if err != nil {
		return err
	}

	if modelTranslation.ReviewerID != "" {
		d.ReviewerID, err = primitive.ObjectIDFromHex(modelTranslation.ReviewerID)
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelTranslation converts bson to model
func (d *Translation) ModelTranslation() *model.Translation {
	Translation := model.Translation{}

	Translation.ID = d.ID.Hex()
	Translation.DatasetID = d.DatasetID.Hex()
	Translation.DatastreamID = d.DatasetID.Hex()
	Translation.ReviewerID = d.ReviewerID.Hex()
	Translation.CreatedAt = d.CreatedAt
	Translation.UpdatedAt = d.UpdatedAt
	Translation.SourceLanguage = d.SourceLanguage
	Translation.SourceSentence = d.SourceSentence
	Translation.DestinationSentence = d.DestinationSentence
	Translation.DestinationLanguage = d.DestinationLanguage
	Translation.Name = d.Name
	Translation.LineNumber = d.LineNumber

	return &Translation
}
