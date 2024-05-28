package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// DeleteResponse provides delete response
type DeleteResponse struct {
	Message   string `json:"message"`
	OK        bool   `json:"ok"`
	ID        string `json:"translation_id"`
	DeletedAt string `json:"deleted_at"`
}

// String provides string repsentation
func (dr *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", dr.Message, dr.OK)
}

// Delete provides dto for translation update
type Delete struct {
	UserID        string `json:"user_id"`
	TranslationID string `json:"translation_id"`
}

// Validate validates translation delete data
func (d *Delete) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid translation delete data", false},
			},
		)
	}
	return nil
}

// FromReader decodes translation delete data from request
func (d *Delete) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid translation update data", false},
			},
		)
	}

	return nil
}
