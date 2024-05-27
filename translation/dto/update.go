package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// Update provides dto for translation update
type Update struct {
	ID                  string `json:"translation_id"`
	SourceSentence      string `json:"source_sentence"`
	SourceLanguage      string `json:"source_language"`
	DestinationSentence string `json:"destination_sentence"`
	DestinationLanguage string `json:"destination_language"`
	Line                int    `json:"line_number"`
	Name                string `json:"name"`
	DatasetID           string `json:"dataset_id"`
	DatastreamID        string `json:"datastream_id"`
	TranslatorID        string `json:"translator_id"`
	ReviewerID          string `json:"reviewer_id"`
}

// Validate validates translation update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid translation update data", false},
			},
		)
	}
	return nil
}

// FromReader decodes translation update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
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
