package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"
	"obyoy-backend/model"
)

// Me stores personal profile related information
type Me struct {
	ID         string      `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Gender     string      `json:"gender"`
	BirthDate  BirthDate   `json:"birth_date"`
	Email      string      `json:"email"`
	Profile    UserDetails `json:"profile"`
	ProfilePic string      `json:"profile_pic"`
	Suspended  bool        `json:"suspended"`
}

// FromModel converts model data to json format data
func (m *Me) FromModel(user *model.User) {
	m.ID = user.ID
	m.FirstName = user.FirstName
	m.LastName = user.LastName
	m.Gender = user.Gender
	m.BirthDate.FromModel(&user.BirthDate)
	m.Email = user.Email
	m.Profile.Day = user.Profile.Day
	m.ProfilePic = user.ProfilePic
	m.Suspended = user.Suspended
}

// MeUpdate stores personal profile update related data
type MeUpdate struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Gender    string      `json:"gender"`
	BirthDate BirthDate   `json:"birth_date"`
	Profile   UserDetails `json:"profile"`
}

// ToModel converts json data to model data
func (mu *MeUpdate) ToModel(user *model.User) {
	user.FirstName = mu.FirstName
	user.LastName = mu.LastName
	user.Gender = mu.Gender
	mu.BirthDate.ToModel(&user.BirthDate)
	user.Profile.Day = mu.Profile.Day
}

// FromReader decodes request data to json type data
func (m *MeUpdate) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(m)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid update data", false},
		})
	}

	return nil
}
