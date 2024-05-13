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

// Deleter provides an interface for updating translations
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes translation
type delete struct {
	storetranslation storetranslation.Translations
	validate         *validator.Validate
}

func (d *delete) toModel(usertranslation *dto.Delete) (translation *model.Translation) {
	translation = &model.Translation{}

	translation.UpdatedAt = time.Now().UTC()
	//	translation.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modeltranslation *model.Translation,
) {
	modeltranslation = d.toModel(delete)
	return
}

func (d *delete) askStore(modeltranslation *model.Translation) (
	id string,
	err error,
) {
	id, err = d.storetranslation.Save(modeltranslation)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Translation,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{}).Debug("User deleted translation successfully")

	return &dto.DeleteResponse{
		Message: "translation deleted",
		OK:      true,
		ID:      id,
		//		DeleteTime: modelNotice.DeletedAt.String(),
	}
}

func (d *delete) giveError() (err error) {
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

// Delete implements Delete interface
func (d *delete) Delete(delete *dto.Delete) (
	*dto.DeleteResponse, error,
) {
	if err := d.validateData(delete); err != nil {
		return nil, err
	}

	modeltranslation := d.convertData(delete)
	id, err := d.askStore(modeltranslation)
	if err == nil {
		return d.giveResponse(modeltranslation, id), nil
	}

	logrus.Error("Could not delete translation ", err)
	err = d.giveError()
	return nil, err
}

// NewDelete returns new instance of NewDelete
func NewDelete(store storetranslation.Translations, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
