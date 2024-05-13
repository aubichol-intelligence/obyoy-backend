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

// Deleter provides an interface for updating datastreams
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes datastream
type delete struct {
	storedatastream storedatastream.Datastreams
	validate        *validator.Validate
}

func (d *delete) toModel(userdatastream *dto.Delete) (datastream *model.Datastream) {
	datastream = &model.Datastream{}

	datastream.UpdatedAt = time.Now().UTC()
	//	datastream.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modeldatastream *model.Datastream,
) {
	modeldatastream = d.toModel(delete)
	return
}

func (d *delete) askStore(modeldatastream *model.Datastream) (
	id string,
	err error,
) {
	id, err = d.storedatastream.Save(modeldatastream)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Datastream,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{}).Debug("User deleted datastream successfully")

	return &dto.DeleteResponse{
		Message: "datastream deleted",
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

	modeldatastream := d.convertData(delete)
	id, err := d.askStore(modeldatastream)
	if err == nil {
		return d.giveResponse(modeldatastream, id), nil
	}

	logrus.Error("Could not delete datastream ", err)
	err = d.giveError()
	return nil, err
}

// NewDelete returns new instance of NewDelete
func NewDelete(store storedatastream.Datastreams, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
