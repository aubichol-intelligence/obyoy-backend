package model

import (
	"encoding/json"
	"time"
)

// Session defines user's session
type Session struct {
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UserType  string    `json:"user_type"`
}

func (s *Session) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Session) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
