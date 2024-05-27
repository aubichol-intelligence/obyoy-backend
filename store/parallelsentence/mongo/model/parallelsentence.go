package model

import (
	"fmt"
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Parallelsentence holds db data type for parallelsentence
type Parallelsentence struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	Name                string               `bson:"name,omitempty"`
	DatasetID           primitive.ObjectID   `bson:"dataset_id"`
	DatastreamID        primitive.ObjectID   `bson:"datastream_id"`
	SourceSentence      string               `bson:"source_sentence"`
	SourceLanguage      string               `bson:"source_language"`
	DestinationSentence string               `bson:"destination_sentence"`
	DestinationLanguage string               `bson:"destination_language"`
	TimesReviewed       int                  `bson:"times_reviewed"`
	TranslatorID        primitive.ObjectID   `bson:"translator_id"`
	Reviewers           []primitive.ObjectID `bson:"reviewers"`
	ReviewedLines       []string             `bson:"reviewed_lines"`
	CreatedAt           time.Time            `bson:"created_at,omitempty"`
	UpdatedAt           time.Time            `bson:"updated_at,omitempty"`
	IsDeleted           bool                 `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for parallelsentences
func (d *Parallelsentence) FromModel(modelParallelsentence *model.Parallelsentence) error {
	d.CreatedAt = modelParallelsentence.CreatedAt
	d.UpdatedAt = modelParallelsentence.UpdatedAt
	d.SourceSentence = modelParallelsentence.SourceSentence
	d.SourceLanguage = modelParallelsentence.SourceLanguage
	d.DestinationSentence = modelParallelsentence.DestinationSentence
	d.DestinationLanguage = modelParallelsentence.DestinationLanguage
	d.TimesReviewed = modelParallelsentence.TimesReviewed
	d.ReviewedLines = modelParallelsentence.ReviewedLines
	d.IsDeleted = modelParallelsentence.IsDeleted

	var err error

	if modelParallelsentence.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelParallelsentence.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	if modelParallelsentence.DatasetID != "" {
		d.DatasetID, err = primitive.ObjectIDFromHex(modelParallelsentence.DatasetID)
	}

	if err != nil {
		return err
	}

	if modelParallelsentence.DatastreamID != "" {
		d.DatastreamID, err = primitive.ObjectIDFromHex(modelParallelsentence.DatastreamID)
	}

	if err != nil {
		return err
	}

	if modelParallelsentence.TranslatorID != "" {
		d.TranslatorID, err = primitive.ObjectIDFromHex(modelParallelsentence.TranslatorID)
	}

	if err != nil {
		return err
	}

	for _, val := range modelParallelsentence.Reviewers {
		id, err := primitive.ObjectIDFromHex(val)
		fmt.Println(err)
		d.Reviewers = append(d.Reviewers, id)
	}

	return nil
}

// ModelParallelsentence converts bson to model
func (d *Parallelsentence) ModelParallelsentence() *model.Parallelsentence {
	Parallelsentence := model.Parallelsentence{}
	Parallelsentence.ID = d.ID.Hex()
	Parallelsentence.CreatedAt = d.CreatedAt
	Parallelsentence.UpdatedAt = d.UpdatedAt
	Parallelsentence.SourceSentence = d.SourceSentence
	Parallelsentence.SourceLanguage = d.SourceLanguage
	Parallelsentence.DestinationSentence = d.DestinationSentence
	Parallelsentence.DestinationLanguage = d.DestinationLanguage
	Parallelsentence.TimesReviewed = d.TimesReviewed

	return &Parallelsentence
}
