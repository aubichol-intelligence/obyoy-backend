package server

import (
	"obyoy-backend/container"
	"obyoy-backend/token"
)

// Token registers token related providers
func Token(c container.Container) {
	c.Register(token.NewRegisterStore)
}
