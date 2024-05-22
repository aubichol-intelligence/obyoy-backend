package model

import (
	"obyoy-backend/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Dataset holds db data type for deliveries
type Dataset struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	Set             []string           `bson:"set,omitempty"`
	UploaderID      primitive.ObjectID `bson:"uploader_id"`
	SourceLanguage  string             `bson:"source_language"`
	TotalLines      int32              `bson:"total_lines"`
	TranslatedLines int32              `bson:"translated_lines"`
	ReviewedLines   int32              `bson:"reviewed_lines"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	IsDeleted       bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Dataset) FromModel(modelDataset *model.Dataset) error {
	d.CreatedAt = modelDataset.CreatedAt
	d.UpdatedAt = modelDataset.UpdatedAt
	d.Set = modelDataset.Set
	d.Name = modelDataset.Name
	d.SourceLanguage = modelDataset.SourceLanguage
	d.TotalLines = modelDataset.TotalLines
	d.TranslatedLines = modelDataset.TranslatedLines
	d.ReviewedLines = modelDataset.ReviewedLines

	var err error

	if modelDataset.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelDataset.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	if modelDataset.UploaderID != "" {
		d.UploaderID, err = primitive.ObjectIDFromHex(modelDataset.UploaderID)
	}

	if err != nil {
		return err
	}

	return nil
}

// ModelDelivery converts bson to model
func (d *Dataset) ModelDataset() *model.Dataset {
	Dataset := model.Dataset{}
	Dataset.ID = d.ID.Hex()
	Dataset.CreatedAt = d.CreatedAt
	Dataset.UpdatedAt = d.UpdatedAt

	return &Dataset
}
