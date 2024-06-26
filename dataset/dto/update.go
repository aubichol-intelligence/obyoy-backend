package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// Update provides dto for dataset update
type Update struct {
	ID              string   `json:"dataset_id"`
	Set             []string `json:"set"`
	Name            string   `json:"name"`
	TotalLines      int32    `json:"total_lines"`
	SourceLanguage  string   `json:"source_language"`
	Description     string   `json:"description"`
	Remarks         string   `json:"remarks"`
	UploaderID      string   `json:"uploader_id"`
	TranslatedLines int32    `json:"translated_lines"`
	ReviewedLines   int32    `json:"reviewed_lines"`
	IsDeleted       bool     `json:"is_deleted"`
}

// Validate validates dataset update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid dataset update data", false},
			},
		)
	}
	return nil
}

// FromReader decodes dataset update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid dataset update data", false},
			},
		)
	}

	return nil
}
