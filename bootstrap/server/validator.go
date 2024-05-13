package server

import (
	"obyoy-backend/container"

	"gopkg.in/go-playground/validator.v9"
)

// Validator registers validation provider
func Validator(c container.Container) {
	c.Register(validator.New)
}
