package datastream

import "obyoy-backend/model"

// Datastreams wraps dataset's functionality
type Datastreams interface {
	Save(*model.Datastream) (id string, err error)
	FindByID(id string) (*model.Datastream, error)
	FindNext() (*model.Datastream, error)
	FindByDatastreamID(id string, skip int64, limit int64) ([]*model.Datastream, error)
	CountByDatastreamID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Datastream, error)
	Search(q string, skip, limit int64) ([]*model.Datastream, error)
	FindByUser(id string, skip, limit int64) ([]*model.Datastream, error)
	FindByDriver(id string) (*model.Datastream, error)
	Count() (int64, error)
}
