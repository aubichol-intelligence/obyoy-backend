package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Translation struct {
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Translation) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Translation) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
