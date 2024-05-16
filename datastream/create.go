package datastream

import (
	"fmt"
	"time"

	"obyoy-backend/datastream/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedatastream "obyoy-backend/store/datastream"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating datastream
type Creater interface {
	Create(create *dto.Create) (*dto.CreateResponse, error)
}

// create creates datastream
type create struct {
	storedatastream storedatastream.Datastreams
	validate        *validator.Validate
}

func (c *create) toModel(userdatastream *dto.Create) (
	datastream *model.Datastream,
) {
	datastream = &model.Datastream{}
	datastream.CreatedAt = time.Now().UTC()
	datastream.UpdatedAt = datastream.CreatedAt
	datastream.ID = userdatastream.ID
	datastream.SourceSentence = userdatastream.SourceSentence
	datastream.LineNumber = userdatastream.LineNumber
	datastream.DatasetID = userdatastream.DatasetID
	datastream.TimesTranslated = userdatastream.TimesTranslated
	datastream.TimesReviewed = userdatastream.TimesReviewed

	return
}

func (c *create) validateData(create *dto.Create) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Create) (
	modeldatastream *model.Datastream,
) {
	modeldatastream = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Datastream) (
	id string,
	err error,
) {
	id, err = c.storedatastream.Save(model)
	return
}

func (c *create) giveResponse(modeldatastream *model.Datastream, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{}).Debug("User created datastream successfully")

	return &dto.CreateResponse{
		Message: "datastream created",
		OK:      true,
		//		datastreamTime: modeldatastream.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create datastream. Error: ", err)
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

	modeldatastream := c.convertData(create)
	id, err := c.askStore(modeldatastream)
	if err == nil {
		return c.giveResponse(modeldatastream, id)
	}

	err = c.giveError()
	return nil, err
}

// CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	Storedatastreams storedatastream.Datastreams
	Validate         *validator.Validate
}

// NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.Storedatastreams,
		params.Validate,
	}
}
