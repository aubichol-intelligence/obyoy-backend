package server

import (
	"obyoy-backend/container"
	dataset "obyoy-backend/store/dataset/mongo"
	datastream "obyoy-backend/store/datastream/mongo"
	parallelsentence "obyoy-backend/store/parallelsentence/mongo"
	translation "obyoy-backend/store/translation/mongo"
	user "obyoy-backend/store/user/mongo"
)

// Store provides constructors for mongo db implementations
func Store(c container.Container) {
	c.Register(dataset.Store)
	c.Register(user.Store)
	c.Register(datastream.Store)
	c.Register(parallelsentence.Store)
	c.Register(translation.Store)
}
