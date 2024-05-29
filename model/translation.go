package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Translation struct {
	ID                  string    `json:"translation_id"`
	UserID              string    `json:"user_id"`
	Name                string    `json:"name"`
	SourceLanguage      string    `json:"source_language"`
	SourceSentence      string    `json:"source_sentence"`
	DestinationLanguage string    `json:"destination_language"`
	DestinationSentence string    `json:"destination_sentence"`
	DatasetID           string    `json:"dataset_id"`
	DatastreamID        string    `json:"datastream_id"`
	LineNumber          int       `json:"line_number"`
	TranslatorID        string    `json:"translator_id"`
	ReviewerID          string    `json:"reviewer_id"`
	ExpiredAt           time.Time `json:"expire_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
	IsDeleted           bool      `json:"is_deleted"`
}

func (s *Translation) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Translation) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
