package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// dataset provides dto for dataset request
type Create struct {
	ID              string   `json:"dataset_id"`
	Set             []string `json:"set"`
	Name            string   `json:"name"`
	TotalLines      int32    `json:"total_lines"`
	SourceLanguage  string   `json:"source_language"`
	UploaderID      string   `json:"uploader_id"`
	TranslatedLines int32    `json:"translated_lines"`
	ReviewedLines   int32    `json:"reviewed_lines"`
	IsDeleted       bool     `json:"is_deleted"`
}

// Validate validates dataset request data
func (d *Create) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for dataset", false},
			},
		)
	}
	return nil
}

// FromReader reads dataset request from request body
func (d *Create) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid dataset data", false},
		})
	}

	return nil
}
