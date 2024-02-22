package cache

import "obyoy-backend/model"

// Session wraps methods to handle session
type Session interface {
	Create(*model.Session) error
	GetByKey(id string) (*model.Session, error)
	GetByUserID(userID string) ([]*model.Session, error)
	RemoveByKey(id string) error
}