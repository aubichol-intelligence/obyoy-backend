package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Datastream struct {
	ID              string    `json:"id"`
	SourceSentence  string    `json:"source_sentence"`
	SourceLanguage  string    `json:"source_language"`
	LineNumber      int32     `json:"line_number"`
	DatasetID       string    `json:"dataset_id"`
	IsTranslated    int32     `json:"is_translated"`
	TimesTranslated int32     `json:"times_translated"`
	TimesReviewed   int32     `json:"times_reviewed"`
	ExpiredAt       time.Time `json:"expire_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (s *Datastream) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Datastream) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
