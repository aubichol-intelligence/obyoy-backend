package server

import (
	"obyoy-backend/container"
	"obyoy-backend/parallelsentence"
)

// Parallelsentence registers parallel sentence related providers
func Parallelsentence(c container.Container) {
	c.Register(parallelsentence.NewCreate)
	c.Register(parallelsentence.NewReader)
	c.Register(parallelsentence.NewUpdate)
	c.Register(parallelsentence.NewDelete)
	c.Register(parallelsentence.NewList)
}
