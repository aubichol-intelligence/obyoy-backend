package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Datastream struct {
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Datastream) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Datastream) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
