package parallelsentence

import (
	"fmt"
	"time"

	"obyoy-backend/errors"
	"obyoy-backend/model"
	"obyoy-backend/parallelsentence/dto"
	storeparallelsentence "obyoy-backend/store/parallelsentence"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Deleter provides an interface for updating parallelsentences
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes parallelsentence
type delete struct {
	storeparallelsentence storeparallelsentence.Parallelsentences
	validate              *validator.Validate
}

func (d *delete) toModel(userparallelsentence *dto.Delete) (parallelsentence *model.Parallelsentence) {
	parallelsentence = &model.Parallelsentence{}

	parallelsentence.ID = userparallelsentence.ParallelsentenceID
	parallelsentence.UpdatedAt = time.Now().UTC()
	parallelsentence.IsDeleted = true

	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelparallelsentence *model.Parallelsentence,
) {
	modelparallelsentence = d.toModel(delete)
	return
}

func (d *delete) askStore(modelparallelsentence *model.Parallelsentence) (
	id string,
	err error,
) {
	id, err = d.storeparallelsentence.Save(modelparallelsentence)
	return
}

func (d *delete) giveResponse(
	modelParallelsentence *model.Parallelsentence,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{}).Debug("User deleted parallelsentence successfully")

	return &dto.DeleteResponse{
		Message:    "parallelsentence deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelParallelsentence.UpdatedAt.String(),
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

	modelparallelsentence := d.convertData(delete)
	id, err := d.askStore(modelparallelsentence)
	if err == nil {
		return d.giveResponse(modelparallelsentence, id), nil
	}

	logrus.Error("Could not delete parallelsentence ", err)
	err = d.giveError()
	return nil, err
}

// NewDelete returns new instance of NewDelete
func NewDelete(store storeparallelsentence.Parallelsentences, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
