package translation

import (
	"fmt"
	"time"

	"obyoy-backend/errors"
	"obyoy-backend/model"
	storetranslation "obyoy-backend/store/translation"
	"obyoy-backend/translation/dto"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Updater provides an interface for updating translations
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates translation
type update struct {
	storetranslation storetranslation.Translations
	validate         *validator.Validate
}

func (u *update) toModel(usertranslation *dto.Update) (translation *model.Translation) {

	translation = &model.Translation{}

	translation.UpdatedAt = time.Now().UTC()
	translation.ID = usertranslation.ID

	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modeltranslation *model.Translation,
) {
	modeltranslation = u.toModel(update)
	return
}

func (u *update) askStore(modeltranslation *model.Translation) (
	id string,
	err error,
) {
	id, err = u.storetranslation.Save(modeltranslation)
	return
}

func (u *update) giveResponse(
	modeltranslation *model.Translation,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		//		"id": modeltranslation.UserID,
	}).Debug("User updated translation successfully")

	return &dto.UpdateResponse{
		Message:    "translation updated",
		OK:         true,
		ID:         id,
		UpdateTime: modeltranslation.UpdatedAt.String(),
	}
}

func (u *update) giveError() (err error) {
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}
	err = fmt.Errorf(
		"%s %w",
		err.Error(),
		&errResp,
	)
	return
}

// Update implements Update interface
func (u *update) Update(update *dto.Update) (
	*dto.UpdateResponse, error,
) {
	if err := u.validateData(update); err != nil {
		return nil, err
	}

	modeltranslation := u.convertData(update)
	id, err := u.askStore(modeltranslation)
	if err == nil {
		return u.giveResponse(modeltranslation, id), nil
	}

	logrus.Error("Could not update translation ", err)
	err = u.giveError()
	return nil, err
}

// NewUpdate returns new instance of NewUpdate
func NewUpdate(store storetranslation.Translations, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
