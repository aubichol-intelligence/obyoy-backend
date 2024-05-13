package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Dataset struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func (s *Dataset) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Dataset) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
