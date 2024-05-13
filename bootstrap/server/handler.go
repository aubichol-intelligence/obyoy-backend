package server

import (
	"obyoy-backend/api"
	"obyoy-backend/container"
)

// Handler registers provider that returns root handler
func Handler(c container.Container) {
	c.Register(api.Handler)
}
