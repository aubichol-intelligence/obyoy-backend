package contest

import (
	"fmt"
	"time"

	"ardent-backend/contest/dto"
	"ardent-backend/errors"
	"ardent-backend/model"
	storecontest "ardent-backend/store/contest"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Deleter provides an interface for updating contests
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes contest
type delete struct {
	storecontest storecontest.Contests
	validate     *validator.Validate
}

func (d *delete) toModel(usercontest *dto.Delete) (contest *model.Contest) {
	contest = &model.Contest{}

	contest.UpdatedAt = time.Now().UTC()
	contest.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelcontest *model.Contest,
) {
	modelcontest = d.toModel(delete)
	return
}

func (d *delete) askStore(modelcontest *model.Contest) (
	id string,
	err error,
) {
	id, err = d.storecontest.Save(modelcontest)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Contest,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{}).Debug("User deleted contest successfully")

	return &dto.DeleteResponse{
		Message: "contest deleted",
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

	modelcontest := d.convertData(delete)
	id, err := d.askStore(modelcontest)
	if err == nil {
		return d.giveResponse(modelcontest, id), nil
	}

	logrus.Error("Could not delete contest ", err)
	err = d.giveError()
	return nil, err
}

// NewDelete returns new instance of NewDelete
func NewDelete(store storecontest.Contests, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
