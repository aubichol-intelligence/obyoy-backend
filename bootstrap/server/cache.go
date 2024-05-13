package server

import (
	"obyoy-backend/cache/redis"
	"obyoy-backend/container"
)

func Cache(c container.Container) {
	c.Register(redis.NewSession)
	//	c.Register(redis.NewConnectionStatus)
}
