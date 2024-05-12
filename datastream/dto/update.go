package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"ardent-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// Update provides dto for contest update
type Update struct {
	ID                string  `json:"contest_id"`
	Name              string  `json:"name"`
	Phone             string  `json:"phone_number"`
	Address           string  `json:"address"`
	UserID            string  `json:"user_id"`
	Duration          int     `json:"duration"`
	RestaurantAddress string  `json:"restaurant_address"`
	RestaurantName    string  `json:"restaurant_name"`
	RestaurantPhone   string  `json:"restaurant_phone"`
	Note              string  `json:"note"`
	Amount            float64 `json:"amount"`
	ItemCount         int     `json:"item_count"`
	State             string  `json:"state"`
}

// Validate validates contest update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid contest update data", false},
			},
		)
	}
	return nil
}

// FromReader decodes contest update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid contest update data", false},
			},
		)
	}

	return nil
}
