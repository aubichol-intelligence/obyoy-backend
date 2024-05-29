package server

import (
	"obyoy-backend/container"
	"obyoy-backend/dataset"
)

// Dataset registers dataset related providers
func Dataset(c container.Container) {
	c.Register(dataset.NewCreate)
	c.Register(dataset.NewReader)
	c.Register(dataset.NewUpdate)
	c.Register(dataset.NewDelete)
	c.Register(dataset.NewList)
	c.Register(dataset.NewCountByStatusReader)
}
