package contest

import "ardent-backend/model"

// Contests wraps delivery's functionality
type Contests interface {
	Save(*model.Contest) (id string, err error)
	FindByID(id string) (*model.Contest, error)
	FindByContestID(id string, skip int64, limit int64) ([]*model.Contest, error)
	CountByContestID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Contest, error)
	Search(q string, skip, limit int64) ([]*model.Contest, error)
	FindByUser(id string, skip, limit int64) ([]*model.Contest, error)
	FindByDriver(id string) (*model.Contest, error)
}
