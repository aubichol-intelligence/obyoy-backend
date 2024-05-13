package translation

import "obyoy-backend/model"

// Translations wraps delivery's functionality
type Translations interface {
	Save(*model.Translation) (id string, err error)
	FindByID(id string) (*model.Translation, error)
	FindByTranslationID(id string, skip int64, limit int64) ([]*model.Translation, error)
	CountByTranslationID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Translation, error)
	Search(q string, skip, limit int64) ([]*model.Translation, error)
	FindByUser(id string, skip, limit int64) ([]*model.Translation, error)
	FindByDriver(id string) (*model.Translation, error)
}
