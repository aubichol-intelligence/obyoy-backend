package model

import (
	"encoding/json"
	"time"
)

// Translation defines user's translation
type Parallelsentence struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"`
	UserID    string    `json:"user_id"`
	ExpiredAt time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json":"updated_at"`
}

func (s *Parallelsentence) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Parallelsentence) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
