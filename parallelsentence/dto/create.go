package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// parallelsentence provides dto for parallelsentence request
type Create struct {
	ID                  string   `json:"parallelsentence_id"`
	SourceSentence      string   `json:"source_sentence"`
	DatasetID           string   `json:"dataset_id"`
	DatastreamID        string   `json:"datastream_id"`
	SourceLanguage      string   `json:"source_language"`
	DestinationSentence string   `json:"destination_sentence"`
	DestinationLanguage string   `json:"destination_language"`
	TimesReviewed       int      `json:"times_reviewed"`
	TranslatorID        string   `json:"translator_id"`
	Reviewers           []string `json:"reviewers"`
	ReviewedLines       []string `json:"reviewed_lines"`
	IsDeleted           bool     `json:"is_deleted"`
}

// Validate validates parallelsentence request data
func (d *Create) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for parallelsentence", false},
			},
		)
	}
	return nil
}

// FromReader reads parallelsentence request from request body
func (d *Create) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid parallelsentence data", false},
		})
	}

	return nil
}
