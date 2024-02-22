package model

import (
	"horkora-backend/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User defines mongodb data type for User
// Account type can be of driver, restaurant and admin
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `bson:"first_name,omitempty"`
	LastName     string             `bson:"last_name,omitempty"`
	Gender       string             `bson:"gender,omitempty"`
	BirthDate    BirthDate          `bson:"birth_date,omitempty"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Verified     bool               `bson:"verified,omitempty"`
	Profile      UserDetails        `bson:"profile,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	Real         bool               `bson:"real,omitempty"`
	ProfilePic   string             `bson:"profile_pic,omitempty"`
	Suspended    bool               `bson:"suspended,omitempty"`
	IsDriver     bool               `bson:"is_driver,omitempty"`
	UserType     string             `bson:"user_type,omitempty"`
	PhoneNumber  string             `bson:"phone_number,omitempty"`
	Address      string             `bson:"address,omitempty"`
	AccountType  string             `bson:"account_type,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty"`
	RestaurantID primitive.ObjectID `bson:"restaurant_id,omitempty"`
}

// TO-DO: need to decide the details later
type UserDetails struct {
	Year  int
	Month int
	Day   int
}

// FromModel converts model data to mongodb model data for user
func (u *User) FromModel(modelUser *model.User) error {
	u.FirstName = modelUser.FirstName
	u.LastName = modelUser.LastName
	u.Gender = modelUser.Gender
	u.BirthDate.FromModel(&modelUser.BirthDate)
	u.Email = modelUser.Email
	u.Password = modelUser.Password
	u.Verified = modelUser.Verified
	u.CreatedAt = modelUser.CreatedAt
	u.UpdatedAt = modelUser.UpdatedAt

	u.Profile.Day = modelUser.Profile.Day
	u.Profile.Month = modelUser.BirthDate.Month
	u.Profile.Year = modelUser.BirthDate.Year
	u.Suspended = modelUser.Suspended
	u.IsDriver = modelUser.IsDriver
	u.ProfilePic = modelUser.ProfilePic
	u.PhoneNumber = modelUser.PhoneNumber
	u.Address = modelUser.Address
	u.AccountType = modelUser.AccountType
	u.Latitude = modelUser.Latitude
	u.Longitude = modelUser.Longitude

	var err error

	if modelUser.RestaurantID != "" {
		u.RestaurantID, err = primitive.ObjectIDFromHex(modelUser.RestaurantID)
	}

	if err != nil {
		return err
	}
	if modelUser.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelUser.ID)
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

// ModelUser converts bson to model for user
func (u *User) ModelUser() *model.User {
	user := model.User{}
	user.ID = u.ID.Hex()
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Gender = u.Gender
	user.BirthDate = *u.BirthDate.ModelBirthDate()
	user.Email = u.Email
	user.Password = u.Password
	user.Verified = u.Verified
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
	user.Profile.Day = u.Profile.Day
	user.Profile.Month = u.BirthDate.Month
	user.Profile.Year = u.BirthDate.Year
	user.Suspended = u.Suspended
	user.IsDriver = u.IsDriver
	user.PhoneNumber = u.PhoneNumber
	user.Address = u.Address
	user.UserType = u.UserType
	user.ProfilePic = u.ProfilePic
	user.Longitude = u.Longitude
	user.Latitude = u.Latitude
	user.AccountType = u.AccountType
	user.RestaurantID = u.RestaurantID.Hex()
	return &user
}
