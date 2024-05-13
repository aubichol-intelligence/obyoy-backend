package server

import (
	"context"
	"time"

	"obyoy-backend/config"
	"obyoy-backend/container"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo provides a constructor to dig component that registers a mongodb client
func Mongo(c container.Container) {
	c.Register(func(cfg config.Mongo) *mongo.Client {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
		if err != nil {
			logrus.Fatal(err)
		}

		if err := client.Ping(context.Background(), nil); err != nil {
			logrus.Fatal(err)
		}

		return client
	})
}

// Providers for all the collections of the database is made.
// RegisterWithName is used since different functions return the same type and it's important
// to differentiate between them
func MongoCollections(c container.Container) {
	registerCollectionProvider(c, "users")
	registerCollectionProvider(c, "tokens")
	registerCollectionProvider(c, "datasets")
	registerCollectionProvider(c, "datastreams")
	registerCollectionProvider(c, "translations")
	registerCollectionProvider(c, "parallelsentences")
}

func registerCollectionProvider(c container.Container, collectionName string) {
	c.RegisterWithName(func(cfg config.Mongo, client *mongo.Client) *mongo.Collection {
		return client.Database(cfg.Database).Collection(collectionName)
	}, collectionName)
}
