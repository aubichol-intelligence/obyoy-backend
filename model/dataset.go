package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Dataset struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Set             []string  `json:"set"`
	TotalLines      int32     `json:"total_lines"`
	SourceLanguage  string    `json:"source_language"`
	UploaderID      string    `json:"uploader_id"`
	TranslatedLines int32     `json:"translated_lines"`
	ReviewedLines   int32     `json:"reviewed_lines"`
	ExpiredAt       time.Time `json:"expire_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsDeleted       bool      `json:"is_deleted"`
}

func (s *Dataset) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Dataset) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
