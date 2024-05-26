package server

import (
	"obyoy-backend/api/dataset"
	"obyoy-backend/api/datastream"
	"obyoy-backend/api/parallelsentence"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/api/translation"
	"obyoy-backend/api/user"

	"obyoy-backend/container"
)

// Route registers all the route providers to container
func Route(c container.Container) {
	c.RegisterGroup(dataset.CreateRoute, "route")
	c.RegisterGroup(dataset.ReadRoute, "route")
	c.RegisterGroup(dataset.UpdateRoute, "route")
	c.RegisterGroup(dataset.DeleteRoute, "route")
	c.RegisterGroup(dataset.ListRoute, "route")
	c.RegisterGroup(routeutils.NewOptionRoute, "route")

	c.RegisterGroup(datastream.CreateRoute, "route")
	c.RegisterGroup(datastream.ReadRoute, "route")
	c.RegisterGroup(datastream.UpdateRoute, "route")
	c.RegisterGroup(datastream.DeleteRoute, "route")
	c.RegisterGroup(datastream.ReadNextRoute, "route")

	c.RegisterGroup(translation.CreateRoute, "route")
	c.RegisterGroup(translation.ReadRoute, "route")
	c.RegisterGroup(translation.UpdateRoute, "route")
	c.RegisterGroup(translation.DeleteRoute, "route")

	c.RegisterGroup(user.RegistrationRoute, "route")
	c.RegisterGroup(user.LoginRoute, "route")
	c.RegisterGroup(user.SearchRoute, "route")
	c.RegisterGroup(user.ListRoute, "route")
	c.RegisterGroup(user.LogoutRoute, "route")

	c.RegisterGroup(parallelsentence.CreateRoute, "route")
	c.RegisterGroup(parallelsentence.UpdateRoute, "route")
	c.RegisterGroup(parallelsentence.DeleteRoute, "route")
	c.RegisterGroup(parallelsentence.ReadRoute, "route")

}
