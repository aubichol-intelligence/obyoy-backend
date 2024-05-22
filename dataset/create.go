package dataset

import (
	"fmt"
	"time"

	"obyoy-backend/dataset/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedataset "obyoy-backend/store/dataset"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating dataset
type Creater interface {
	Create(create *dto.Create) (*dto.CreateResponse, error)
}

// create creates dataset
type create struct {
	storedataset storedataset.Datasets
	validate     *validator.Validate
}

func (c *create) toModel(userdataset *dto.Create) (
	dataset *model.Dataset,
) {
	dataset = &model.Dataset{}
	dataset.CreatedAt = time.Now().UTC()
	dataset.UpdatedAt = dataset.CreatedAt
	dataset.Name = userdataset.Name
	dataset.Set = userdataset.Set
	dataset.SourceLanguage = userdataset.SourceLanguage
	dataset.ReviewedLines = userdataset.ReviewedLines
	dataset.UploaderID = userdataset.UploaderID

	return
}

func (c *create) validateData(create *dto.Create) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Create) (
	modeldataset *model.Dataset,
) {
	modeldataset = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Dataset) (
	id string,
	err error,
) {
	id, err = c.storedataset.Save(model)
	return
}

func (c *create) giveResponse(modeldataset *model.Dataset, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{}).Debug("User created dataset successfully")

	return &dto.CreateResponse{
		Message: "dataset created",
		OK:      true,
		//		datasetTime: modeldataset.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create dataset. Error: ", err)
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

	modeldataset := c.convertData(create)
	id, err := c.askStore(modeldataset)

	if err == nil {
		return c.giveResponse(modeldataset, id)
	}

	err = c.giveError()
	return nil, err
}

// CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	Storedatasets storedataset.Datasets
	Validate      *validator.Validate
}

// NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.Storedatasets,
		params.Validate,
	}
}
