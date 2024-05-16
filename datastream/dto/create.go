package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// contest provides dto for contest request
type Create struct {
	ID              string `json:"datastream_id"`
	SourceSentence  string `json:"source_sentence"`
	LineNumber      int32  `json:"line_number"`
	DatasetID       string `json:"dataset_id"`
	TimesTranslated int32  `json:"times_translated"`
	TimesReviewed   int32  `json:"times_reviewed"`
	Standings       string `json:"standings"`
	LandingURL      string `json:"landing_url"`
	Name            string `json:"name"`
	IsDeleted       bool   `json:"is_deleted"`
}

// Validate validates contest request data
func (d *Create) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for contest", false},
			},
		)
	}
	return nil
}

// FromReader reads contest request from request body
func (d *Create) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid contest data", false},
		})
	}

	return nil
}
