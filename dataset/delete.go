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

// Deleter provides an interface for deleting datasets
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes dataset
type delete struct {
	storedataset storedataset.Datasets
	validate     *validator.Validate
}

func (d *delete) toModel(userdataset *dto.Delete) (dataset *model.Dataset) {
	dataset = &model.Dataset{}

	dataset.UpdatedAt = time.Now().UTC()
	dataset.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modeldataset *model.Dataset,
) {
	modeldataset = d.toModel(delete)
	return
}

func (d *delete) askStore(modeldataset *model.Dataset) (
	id string,
	err error,
) {
	id, err = d.storedataset.Save(modeldataset)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Dataset,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{}).Debug("User deleted dataset successfully")

	return &dto.DeleteResponse{
		Message:    "dataset deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelNotice.UpdatedAt.String(),
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

	modeldataset := d.convertData(delete)
	id, err := d.askStore(modeldataset)
	if err == nil {
		return d.giveResponse(modeldataset, id), nil
	}

	logrus.Error("Could not delete dataset ", err)
	err = d.giveError()
	return nil, err
}

// NewDelete returns new instance of NewDelete
func NewDelete(store storedataset.Datasets, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
