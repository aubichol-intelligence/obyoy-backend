package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"ardent-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// DeleteResponse provides create response
type DeleteResponse struct {
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"contest_id"`
	RequestTime string `json:"request_time"`
}

// String provides string repsentation
func (dr *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", dr.Message, dr.OK)
}

// Delete provides dto for contest update
type Delete struct {
	UserID    string `json:"user_id"`
	ContestID string `json:"contest_id"`
}

// Validate validates contest delete data
func (d *Delete) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid comment update data", false},
			},
		)
	}
	return nil
}

// FromReader decodes contest delete data from request
func (d *Delete) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid comment update data", false},
			},
		)
	}

	return nil
}
