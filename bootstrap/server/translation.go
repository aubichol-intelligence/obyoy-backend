package server

import (
	"obyoy-backend/container"
	"obyoy-backend/translation"
)

// Order registers translation related providers
func Translation(c container.Container) {
	c.Register(translation.NewCreate)
	c.Register(translation.NewReader)
	c.Register(translation.NewUpdate)
	c.Register(translation.NewDelete)
}
