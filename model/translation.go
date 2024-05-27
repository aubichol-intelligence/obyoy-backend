package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Translation struct {
	ID        string    `json:translation_id"`
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func (s *Translation) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Translation) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
