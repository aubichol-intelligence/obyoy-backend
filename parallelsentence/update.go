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

// Updater provides an interface for updating parallelsentences
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates parallelsentence
type update struct {
	storeparallelsentence storeparallelsentence.Parallelsentences
	validate              *validator.Validate
}

func (u *update) toModel(userparallelsentence *dto.Update) (parallelsentence *model.Parallelsentence) {

	parallelsentence = &model.Parallelsentence{}

	parallelsentence.UpdatedAt = time.Now().UTC()
	parallelsentence.ID = userparallelsentence.ID
	parallelsentence.TimesReviewed = userparallelsentence.TimesReviewed
	parallelsentence.SourceSentence = userparallelsentence.SourceSentence
	parallelsentence.SourceLanguage = userparallelsentence.SourceLanguage
	parallelsentence.DestinationSentence = userparallelsentence.DestinationSentence
	parallelsentence.DestinationLanguage = userparallelsentence.DestinationLanguage
	parallelsentence.TranslatorID = userparallelsentence.TranslatorID
	parallelsentence.Reviewers = userparallelsentence.Reviewers
	parallelsentence.ReviewedLines = userparallelsentence.ReviewedLines
	parallelsentence.DatastreamID = userparallelsentence.DatastreamID
	parallelsentence.DatasetID = userparallelsentence.DatasetID

	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelparallelsentence *model.Parallelsentence,
) {
	modelparallelsentence = u.toModel(update)
	return
}

func (u *update) askStore(modelparallelsentence *model.Parallelsentence) (
	id string,
	err error,
) {
	id, err = u.storeparallelsentence.Save(modelparallelsentence)
	return
}

func (u *update) giveResponse(
	modelparallelsentence *model.Parallelsentence,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelparallelsentence.ID,
	}).Debug("User updated parallelsentence successfully")

	return &dto.UpdateResponse{
		Message:    "parallelsentence updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelparallelsentence.UpdatedAt.String(),
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

	modelparallelsentence := u.convertData(update)
	id, err := u.askStore(modelparallelsentence)
	if err == nil {
		return u.giveResponse(modelparallelsentence, id), nil
	}

	logrus.Error("Could not update parallelsentence ", err)
	err = u.giveError()
	return nil, err
}

// NewUpdate returns new instance of NewUpdate
func NewUpdate(store storeparallelsentence.Parallelsentences, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
