package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// datastream provides dto for datastream request
type Create struct {
	ID              string `json:"datastream_id"`
	SourceSentence  string `json:"source_sentence"`
	LineNumber      int32  `json:"line_number"`
	DatasetID       string `json:"dataset_id"`
	TimesTranslated int32  `json:"times_translated"`
	TimesReviewed   int32  `json:"times_reviewed"`
	IsTranslated    int32  `json:"is_tranlated"`
	Name            string `json:"name"`
	IsDeleted       bool   `json:"is_deleted"`
}

// Validate validates datastream request data
func (d *Create) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for datastream", false},
			},
		)
	}
	return nil
}

// FromReader reads datastream request from request body
func (d *Create) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid datastream data", false},
		})
	}

	return nil
}
