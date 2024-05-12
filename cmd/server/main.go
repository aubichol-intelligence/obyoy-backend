package main

import (
	"net/http"
	"obyoy-backend/bootstrap/server"
	"obyoy-backend/config"
	"obyoy-backend/container"

	"github.com/sirupsen/logrus"
)

func main() {

	if err := server.Viper(); err != nil {
		logrus.Fatal(err)
	}

	c := container.New()
	server.Logrus()
	//Cfg provides configuration related providers to the container
	server.Cfg(c)

	server.Handler(c)
	server.Route(c)
	server.Store(c)
	server.Mongo(c)
	server.MongoCollections(c)
	server.Redis(c)
	server.Cache(c)
	server.Middleware(c)
	server.User(c)
	server.WS(c)
	server.Validator(c)
	server.Event(c)

	c.Resolve(
		func(cfg config.Server, handler http.Handler) {
			logrus.Info("Starting server at port ", cfg.Port)
			http.ListenAndServe(":"+cfg.Port, handler)
		},
	)
}
