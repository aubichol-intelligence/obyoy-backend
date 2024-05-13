package server

import (
	"obyoy-backend/config"
	"obyoy-backend/container"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// Redis provides a constructor for creating redis.Client instance from cfg.Redis config details
func Redis(c container.Container) {
	c.Register(func(cfg config.Redis) *redis.Client {
		client := redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB, // use default DB
		})

		if _, err := client.Ping().Result(); err != nil {
			logrus.Fatal(err)
		}

		return client
	})
}
