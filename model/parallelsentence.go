package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Parallelsentence struct {
	ID                  string    `json:"id"`
	Status              string    `json:"status"`
	DatastreamID        string    `json:"datastream_id"`
	DatasetID           string    `json:"dataset_id"`
	DatasetName         string    `json:"name"`
	LineNumber          int       `json:"line_number"`
	SourceSentence      string    `json:"source_sentence"`
	SourceLanguage      string    `json:"source_language"`
	DestinationSentence string    `json:"destination_sentence"`
	DestinationLanguage string    `json:"destination_language"`
	TranslatorID        string    `json:"translator_id"`
	Reviewers           []string  `json:"reviewers"`
	ReviewedLines       []string  `json:"reviewed_lines"`
	TimesReviewed       int       `json:"times_reviewed"`
	ExpiredAt           time.Time `json:"expire_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
	IsDeleted           bool      `json:"is_deleted"`
}

func (s *Parallelsentence) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Parallelsentence) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
