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
}

func (s *Session) ToByte() ([]byte, error) {
	return json.Marshal(s)
}
	//Username  string

func (s *Session) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, s)
}
(base) nelson@NELSONs-MacBook-Pro model % ls
delivery.go		maps.go			order.go		session.go		state.go		token.go
email.go		menuitem.go		restaurant.go		sms.go			staticcontent.go	user.go
(base) nelson@NELSONs-MacBook-Pro model % cat token.go
package model

import "time"

// Token defines token model
type Token struct {
	ID        string
	UserID    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
