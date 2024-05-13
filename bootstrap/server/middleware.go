package server

import (
	"obyoy-backend/api/middleware"
	"obyoy-backend/container"
)

// Middleware registers middleware related providers
func Middleware(c container.Container) {
	c.Register(middleware.NewAuthMiddleware)
	c.Register(middleware.NewAuthMiddlewareURL) // don't know what it does
	c.Register(middleware.MessageNotificationMiddleware)
}
