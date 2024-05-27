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
	TranslatorID        primitive.ObjectID `bson:"translator_id"`
	ReviewerID          primitive.ObjectID `bson:"reviewer_id"`
	CreatedAt           time.Time          `bson:"created_at,omitempty"`
	UpdatedAt           time.Time          `bson:"updated_at,omitempty"`
	SourceSentence      string             `bson:"source_sentence,omitempty"`
	SourceLanguage      string             `bson:"source_language,omitempty"`
	DestinationSentence string             `bson:"destination_sentence"`
	DestinationLanguage string             `bson:"destination_language"`
	Name                string             `bson:"name,omitempty"`
	LineNumber          int                `bson:"line_number,omitempty"`
	IsDeleted           bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Translation) FromModel(modelDelivery *model.Translation) error {
	d.CreatedAt = modelDelivery.CreatedAt
	d.UpdatedAt = modelDelivery.UpdatedAt
	d.SourceLanguage = modelDelivery.SourceLanguage
	d.SourceSentence = modelDelivery.SourceSentence
	d.DestinationSentence = modelDelivery.DestinationSentence
	d.DestinationLanguage = modelDelivery.DestinationLanguage
	d.Name = modelDelivery.Name
	d.LineNumber = modelDelivery.LineNumber

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

	if modelDelivery.DatastreamID != "" {
		d.DatastreamID, err = primitive.ObjectIDFromHex(modelDelivery.DatastreamID)
	}

	if err != nil {
		return err
	}

	if modelDelivery.TranslatorID != "" {
		d.TranslatorID, err = primitive.ObjectIDFromHex(modelDelivery.TranslatorID)
	}

	if err != nil {
		return err
	}

	if modelDelivery.ReviewerID != "" {
		d.ReviewerID, err = primitive.ObjectIDFromHex(modelDelivery.ReviewerID)
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelDelivery converts bson to model
func (d *Translation) ModelTranslation() *model.Translation {
	Translation := model.Translation{}
	Translation.ID = d.ID.Hex()
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
