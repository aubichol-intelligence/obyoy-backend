package translation

import (
	"fmt"
	"time"

	"obyoy-backend/errors"
	"obyoy-backend/model"
	storetranslation "obyoy-backend/store/translation"
	"obyoy-backend/translation/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating translation
type Creater interface {
	Create(create *dto.Create) (*dto.CreateResponse, error)
}

// create creates translation
type create struct {
	storetranslation storetranslation.Translations
	validate         *validator.Validate
}

func (c *create) toModel(usertranslation *dto.Create) (
	translation *model.Translation,
) {
	translation = &model.Translation{}
	translation.CreatedAt = time.Now().UTC()
	translation.UpdatedAt = translation.CreatedAt
	translation.ID = usertranslation.ID
	translation.SourceSentence = usertranslation.SourceSentence
	translation.SourceLanguage = usertranslation.SourceLanguage
	translation.DestinationSentence = usertranslation.DestinationSentence
	translation.DestinationLanguage = usertranslation.DestinationLanguage
	translation.LineNumber = usertranslation.Line
	translation.DatasetID = usertranslation.DatasetID
	translation.DatastreamID = usertranslation.DatastreamID
	translation.ReviewerID = usertranslation.ReviewerID
	translation.Name = usertranslation.Name

	return
}

func (c *create) validateData(create *dto.Create) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Create) (
	modeltranslation *model.Translation,
) {
	modeltranslation = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Translation) (
	id string,
	err error,
) {
	id, err = c.storetranslation.Save(model)
	return
}

func (c *create) giveResponse(modeltranslation *model.Translation, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{}).Debug("User created translation successfully")

	return &dto.CreateResponse{
		Message: "translation created",
		OK:      true,
		//		translationTime: modeltranslation.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create translation. Error: ", err)
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

	modeltranslation := c.convertData(create)
	id, err := c.askStore(modeltranslation)
	if err == nil {
		return c.giveResponse(modeltranslation, id)
	}

	err = c.giveError()
	return nil, err
}

// CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	Storetranslations storetranslation.Translations
	Validate          *validator.Validate
}

// NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.Storetranslations,
		params.Validate,
	}
}
