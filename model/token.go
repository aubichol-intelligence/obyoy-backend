package model

import (
	"encoding/json"
	"time"
)

// Token defines token model
type Token struct {
	ID        string
	UserID    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Token) ToByte() ([]byte, error) {
	return json.Marshal(s)
}

//Username  string

func (s *Token) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
