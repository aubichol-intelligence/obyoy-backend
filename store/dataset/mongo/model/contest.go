package model

import (
	"obyoy-backend/model"
	"time"

	"ardent-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contest holds db data type for deliveries
type Dataset struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name,omitempty"`
	Phone             string             `bson:"phone_number,omitempty"`
	Address           string             `bson:"address,omitempty"`
	UserID            primitive.ObjectID `bson:"user_id,omitempty"`
	DriverID          primitive.ObjectID `bson:"driver_id,omitempty"`
	OrderID           primitive.ObjectID `bson:"order_id,omitempty"`
	RestaurantID      primitive.ObjectID `bson:"restaurant_id,omitempty"`
	RestaurantName    string             `bson:"restaurant_name,omitempty"`
	RestaurantAddress string             `bson:"restaurant_address,omitempty"`
	RestaurantPhone   string             `bson:"restaurant_phone,omitempty"`
	CustomerDetails   Credentials        `bson:"customer_details,omitempty"`
	RestaurantDetails Credentials        `bson:"restaurant_details,omitempty"`
	Note              string             `bson:"note,omitempty,omitempty"`
	Amount            float64            `bson:"amount,omitempty"`
	IsActive          bool               `bson:"is_active,omitempty"`
	Distance          string             `bson:"distance,omitempty"`
	State             string             `bson:"state,omitempty"`
	ItemCount         int                `bson:"item_count,omitempty"`
	CreatedAt         time.Time          `bson:"created_at,omitempty"`
	UpdatedAt         time.Time          `bson:"updated_at,omitempty"`
	IsDeleted         bool               `bson:"is_deleted,omitempty"`
}

type Credentials struct {
	Name      string  `bson:"name,omitempty"`
	Phone     string  `bson:"phone,omitempty"`
	Address   string  `bson:"address,omitempty"`
	Latitude  float64 `bson:"latitude,omitempty"`
	Longitude float64 `bson:"longitude,omitempty"`
}

// FromModel converts model data to db data for deliveries
func (d *Contest) FromModel(modelDelivery *model.Contest) error {
	d.CreatedAt = modelDelivery.CreatedAt
	d.UpdatedAt = modelDelivery.UpdatedAt
	d.Note = modelDelivery.Note

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
func (d *Contest) ModelDataset() *model.Contest {
	Contest := model.Dataset{}
	Contest.ID = d.ID.Hex()
	Contest.CreatedAt = d.CreatedAt
	Contest.UpdatedAt = d.UpdatedAt

	Contest.Note = d.Note

	return &Contest
}
