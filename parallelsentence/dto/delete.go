package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// DeleteResponse provides create response
type DeleteResponse struct {
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"parallelsentence_id"`
	RequestTime string `json:"request_time"`
	DeleteTime  string `json:"delete_time"`
}

// String provides string repsentation
func (dr *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", dr.Message, dr.OK)
}

// Delete provides dto for parallelsentence update
type Delete struct {
	UserID             string `json:"user_id"`
	ParallelsentenceID string `json:"parallelsentence_id"`
}

// Validate validates parallelsentence delete data
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

// FromReader decodes parallelsentence delete data from request
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
