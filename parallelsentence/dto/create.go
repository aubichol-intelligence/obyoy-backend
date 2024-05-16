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
	ID                  string   `json:"contest_id"`
	SourceSentence      string   `json:"source_sentence"`
	SourceLanguage      string   `json:"source_language"`
	DestinationSentence string   `json:"destination_sentence"`
	DestinationLanguage string   `json:"destination_language"`
	TranslatorID        string   `json:"translator_id"`
	Reviewers           []string `json:"reviewers"`
	ReviewedLines       []string `json:"reviewed_lines"`
	ImageURL            string   `json:"image_url"`
	Standings           string   `json:"standings"`
	LandingURL          string   `json:"landing_url"`
	Name                string   `json:"name"`
	IsDeleted           bool     `json:"is_deleted"`
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
