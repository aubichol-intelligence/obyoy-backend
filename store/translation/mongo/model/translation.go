package model

import (
	"time"

	"obyoy-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Translation holds db data type for deliveries
type Translation struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	IsDeleted bool               `bson:"is_deleted,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Translation) FromModel(modelDelivery *model.Translation) error {
	d.CreatedAt = modelDelivery.CreatedAt
	d.UpdatedAt = modelDelivery.UpdatedAt

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
func (d *Translation) ModelTranslation() *model.Translation {
	Translation := model.Translation{}
	Translation.ID = d.ID.Hex()
	Translation.CreatedAt = d.CreatedAt
	Translation.UpdatedAt = d.UpdatedAt

	return &Translation
}
