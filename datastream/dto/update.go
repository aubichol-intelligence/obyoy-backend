package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// Update provides dto for contest update
type Update struct {
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
