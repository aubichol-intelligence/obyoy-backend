package datastream

import (
	"fmt"
	"time"

	"obyoy-backend/datastream/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedatastream "obyoy-backend/store/datastream"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Updater provides an interface for updating datastreams
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates datastream
type update struct {
	storedatastream storedatastream.Datastreams
	validate        *validator.Validate
}

func (u *update) toModel(userdatastream *dto.Update) (datastream *model.Datastream) {

	datastream = &model.Datastream{}

	datastream.UpdatedAt = time.Now().UTC()
	datastream.ID = userdatastream.ID
	datastream.Note = userdatastream.Note

	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modeldatastream *model.Datastream,
) {
	modeldatastream = u.toModel(update)
	return
}

func (u *update) askStore(modeldatastream *model.Datastream) (
	id string,
	err error,
) {
	id, err = u.storedatastream.Save(modeldatastream)
	return
}

func (u *update) giveResponse(
	modeldatastream *model.Datastream,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		//		"id": modeldatastream.UserID,
	}).Debug("User updated datastream successfully")

	return &dto.UpdateResponse{
		Message:    "datastream updated",
		OK:         true,
		ID:         id,
		UpdateTime: modeldatastream.UpdatedAt.String(),
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

	modeldatastream := u.convertData(update)
	id, err := u.askStore(modeldatastream)
	if err == nil {
		return u.giveResponse(modeldatastream, id), nil
	}

	logrus.Error("Could not update datastream ", err)
	err = u.giveError()
	return nil, err
}

// NewUpdate returns new instance of NewUpdate
func NewUpdate(store storedatastream.Datastreams, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
