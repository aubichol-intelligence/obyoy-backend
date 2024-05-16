package parallelsentence

import (
	"fmt"
	"time"

	"obyoy-backend/errors"
	"obyoy-backend/model"
	"obyoy-backend/parallelsentence/dto"
	storeparallelsentence "obyoy-backend/store/parallelsentence"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating parallelsentence
type Creater interface {
	Create(create *dto.Create) (*dto.CreateResponse, error)
}

// create creates parallelsentence
type create struct {
	storeparallelsentence storeparallelsentence.Parallelsentences
	validate              *validator.Validate
}

func (c *create) toModel(userparallelsentence *dto.Create) (
	parallelsentence *model.Parallelsentence,
) {
	parallelsentence = &model.Parallelsentence{}
	parallelsentence.CreatedAt = time.Now().UTC()
	parallelsentence.UpdatedAt = parallelsentence.CreatedAt
	parallelsentence.ID = userparallelsentence.ID
	parallelsentence.SourceSentence = userparallelsentence.SourceSentence
	parallelsentence.SourceLanguage = userparallelsentence.SourceLanguage
	parallelsentence.DestinationSentence = userparallelsentence.DestinationSentence
	parallelsentence.DestinationLanguage = userparallelsentence.DestinationLanguage
	parallelsentence.TranslatorID = userparallelsentence.TranslatorID
	parallelsentence.Reviewers = userparallelsentence.Reviewers
	parallelsentence.ReviewedLines = userparallelsentence.ReviewedLines

	return
}

func (c *create) validateData(create *dto.Create) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Create) (
	modelparallelsentence *model.Parallelsentence,
) {
	modelparallelsentence = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Parallelsentence) (
	id string,
	err error,
) {
	id, err = c.storeparallelsentence.Save(model)
	return
}

func (c *create) giveResponse(modelparallelsentence *model.Parallelsentence, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{}).Debug("User created parallelsentence successfully")

	return &dto.CreateResponse{
		Message: "parallelsentence created",
		OK:      true,
		//		parallelsentenceTime: modelparallelsentence.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create parallelsentence. Error: ", err)
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}

	err = fmt.Errorf("%s %w", err.Error(), &errResp)
	return
}

// Create implements Creater interface
func (c *create) Create(create *dto.Create) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelparallelsentence := c.convertData(create)
	id, err := c.askStore(modelparallelsentence)
	if err == nil {
		return c.giveResponse(modelparallelsentence, id)
	}

	err = c.giveError()
	return nil, err
}

// CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	Storeparallelsentences storeparallelsentence.Parallelsentences
	Validate               *validator.Validate
}

// NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.Storeparallelsentences,
		params.Validate,
	}
}
