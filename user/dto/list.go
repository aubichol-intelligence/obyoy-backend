package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// user provides dto for user request
type ListByType struct {
	State string `json:"account_type"`
	Skip  int64  `json:"skip"`
	Limit int64  `json:"limit"`
}

// Validate validates user data
func (d *ListByType) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for user", false},
			},
		)
	}
	return nil
}

// FromReader reads user from request body
func (d *ListByType) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid user list read data", false},
		})
	}

	return nil
}
