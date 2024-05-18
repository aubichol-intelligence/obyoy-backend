package server

import (
	"obyoy-backend/container"
	"obyoy-backend/datastream"
)

// Order registers order related providers
func Datastream(c container.Container) {
	c.Register(datastream.NewCreate)
	c.Register(datastream.NewReader)
	c.Register(datastream.NewUpdate)
	c.Register(datastream.NewDelete)
	c.Register(datastream.NewNextReader)
}
