package dataset

import (
	"fmt"
	"time"

	"obyoy-backend/dataset/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedataset "obyoy-backend/store/dataset"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Updater provides an interface for updating datasets
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates dataset
type update struct {
	storedataset storedataset.Datasets
	validate     *validator.Validate
}

func (u *update) toModel(userdataset *dto.Update) (dataset *model.Dataset) {

	dataset = &model.Dataset{}

	dataset.UpdatedAt = time.Now().UTC()
	dataset.ID = userdataset.ID
	dataset.Note = userdataset.Note

	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modeldataset *model.Dataset,
) {
	modeldataset = u.toModel(update)
	return
}

func (u *update) askStore(modeldataset *model.Dataset) (
	id string,
	err error,
) {
	id, err = u.storedataset.Save(modeldataset)
	return
}

func (u *update) giveResponse(
	modeldataset *model.Dataset,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		//		"id": modeldataset.UserID,
	}).Debug("User updated dataset successfully")

	return &dto.UpdateResponse{
		Message:    "dataset updated",
		OK:         true,
		ID:         id,
		UpdateTime: modeldataset.UpdatedAt.String(),
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

	modeldataset := u.convertData(update)
	id, err := u.askStore(modeldataset)
	if err == nil {
		return u.giveResponse(modeldataset, id), nil
	}

	logrus.Error("Could not update dataset ", err)
	err = u.giveError()
	return nil, err
}

// NewUpdate returns new instance of NewUpdate
func NewUpdate(store storedataset.Datasets, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
