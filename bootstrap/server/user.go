package server

import (
	"obyoy-backend/container"
	"obyoy-backend/user"
)

// User registers user related providers
func User(c container.Container) {
	c.Register(user.NewRegistry)
	c.Register(user.NewSearcher)
	c.Register(user.NewLogin)
	c.Register(user.NewEmailAndPasswordChecker)
	c.Register(user.NewSessionVerifier)
	c.Register(user.NewMyProfile)
	c.Register(user.NewListSuspender)
	c.Register(user.NewLister)
}
