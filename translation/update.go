package contest

import (
	"fmt"
	"time"

	"obyoy-backend/contest/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storecontest "obyoy-backend/store/contest"

	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Updater provides an interface for updating contests
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates contest
type update struct {
	storecontest storecontest.Contests
	validate     *validator.Validate
}

func (u *update) toModel(usercontest *dto.Update) (contest *model.Contest) {

	contest = &model.Contest{}

	contest.UpdatedAt = time.Now().UTC()
	contest.ID = usercontest.ID
	contest.Note = usercontest.Note

	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelcontest *model.Contest,
) {
	modelcontest = u.toModel(update)
	return
}

func (u *update) askStore(modelcontest *model.Contest) (
	id string,
	err error,
) {
	id, err = u.storecontest.Save(modelcontest)
	return
}

func (u *update) giveResponse(
	modelcontest *model.Contest,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		//		"id": modelcontest.UserID,
	}).Debug("User updated contest successfully")

	return &dto.UpdateResponse{
		Message:    "contest updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelcontest.UpdatedAt.String(),
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

	modelcontest := u.convertData(update)
	id, err := u.askStore(modelcontest)
	if err == nil {
		return u.giveResponse(modelcontest, id), nil
	}

	logrus.Error("Could not update contest ", err)
	err = u.giveError()
	return nil, err
}

// NewUpdate returns new instance of NewUpdate
func NewUpdate(store storecontest.Contests, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
