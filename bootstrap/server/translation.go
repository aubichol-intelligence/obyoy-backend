package server

import (
	"obyoy-backend/container"
	"obyoy-backend/translation"
)

// Translation registers translation related providers
func Translation(c container.Container) {
	c.Register(translation.NewCreate)
	c.Register(translation.NewReader)
	c.Register(translation.NewUpdate)
	c.Register(translation.NewDelete)
	c.Register(translation.NewCountByStatusReader)
}
