package server

import (
	"obyoy-backend/config"
	"obyoy-backend/container"
)

// Cfg registers configuration related providers
func Cfg(c container.Container) {
	//LoadSession provides a constructor to dig container that loads the duration of a session and the maximum number of sessions a user can have
	c.Register(config.LoadSession)
	//LoadMongo provides a container that loads mongodb server url name and the name of the database
	c.Register(config.LoadMongo)
	//LoadRedis provides a contructor to dig container for loading redis related configurations
	c.Register(config.LoadRedis)
	//LoadServer provides a consturctor to dig container for loading go server related configurations
	c.Register(config.LoadServer)
	//LoadConnectionCache provides a constructor to dig container for connection cache
	//TODO: What is connection cache?
	c.Register(config.LoadConnectionCache)
}
