package dataset

import "obyoy-backend/model"

// Datasets wraps dataset's functionality
type Datasets interface {
	Save(*model.Dataset) (id string, err error)
	FindByID(id string) (*model.Dataset, error)
	FindByDatasetID(id string, skip int64, limit int64) ([]*model.Dataset, error)
	CountByDatasetID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Dataset, error)
	Search(q string, skip, limit int64) ([]*model.Dataset, error)
	FindByUser(id string, skip, limit int64) ([]*model.Dataset, error)
	List(skip, limit int64) ([]*model.Dataset, error)
}
