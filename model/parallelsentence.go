package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Parallelsentence struct {
	ID                  string    `json:"id"`
	SourceSentence      string    `json:"source_sentence"`
	SourceLanguage      string    `json:"source_language"`
	DestinationSentence string    `json:"destination_sentence"`
	DestinationLanguage string    `json:"destination_language"`
	TranslatorID        string    `json:"translator_id"`
	Reviewers           []string  `json:"reviewers"`
	ReviewedLines       []string  `json:"reviewed_lines"`
	ExpiredAt           time.Time `json:"expire_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (s *Parallelsentence) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Parallelsentence) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
