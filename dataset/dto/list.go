package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// ReadReq stores order read request data
type ListReq struct {
	UserID  string
	Contest string
	Skip    int64 `json:"skip"`
	Limit   int64 `json:"limit"`
}

// Validate validates restaurant data
func (d *ListReq) Validate(validate *validator.Validate) error {
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

// FromReader reads restaurant from request body
func (d *ListReq) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid user list read data", false},
		})
	}

	return nil
}
