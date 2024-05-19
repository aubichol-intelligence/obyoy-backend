package model

import (
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Parallelsentence holds db data type for deliveries
type Parallelsentence struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	Name                string               `bson:"name,omitempty"`
	SourceSentence      string               `bson:"source_sentence"`
	SourceLanguage      string               `bson:"source_language"`
	DestinationSentence string               `bson:"destination_sentence"`
	DestinationLanguage string               `bson:"destination_language"`
	TranslatorID        primitive.ObjectID   `bson:"translator_id"`
	Reviewers           []primitive.ObjectID `bson:"reviewers"`
	ReviewedLines       []string             `bson:"reviewed_lines"`
	CreatedAt           time.Time            `bson:"created_at,omitempty"`
	UpdatedAt           time.Time            `bson:"updated_at,omitempty"`
	IsDeleted           bool                 `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Parallelsentence) FromModel(modelDelivery *model.Parallelsentence) error {
	d.CreatedAt = modelDelivery.CreatedAt
	d.UpdatedAt = modelDelivery.UpdatedAt
	d.SourceSentence = modelDelivery.SourceSentence
	d.SourceLanguage = modelDelivery.SourceLanguage
	d.DestinationSentence = modelDelivery.DestinationSentence
	d.DestinationLanguage = modelDelivery.DestinationLanguage

	var err error

	if modelDelivery.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelDelivery.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelDelivery converts bson to model
func (d *Parallelsentence) ModelParallelsentence() *model.Parallelsentence {
	Parallelsentence := model.Parallelsentence{}
	Parallelsentence.ID = d.ID.Hex()
	Parallelsentence.CreatedAt = d.CreatedAt
	Parallelsentence.UpdatedAt = d.UpdatedAt

	return &Parallelsentence
}
