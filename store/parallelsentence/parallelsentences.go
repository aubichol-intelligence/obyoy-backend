package parallelsentence

import "obyoy-backend/model"

// Parallelsentences wraps parallelsentence's functionality
type Parallelsentences interface {
	Save(*model.Parallelsentence) (id string, err error)
	FindByID(id string) (*model.Parallelsentence, error)
	FindByParallelsentenceID(id string, skip int64, limit int64) ([]*model.Parallelsentence, error)
	CountByParallelsentenceID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Parallelsentence, error)
	Search(q string, skip, limit int64) ([]*model.Parallelsentence, error)
	FindByUser(id string, skip, limit int64) ([]*model.Parallelsentence, error)
	FindByDriver(id string) (*model.Parallelsentence, error)
	List(skip, limit int64) ([]*model.Parallelsentence, error)
}
