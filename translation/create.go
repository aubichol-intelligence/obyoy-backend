package contest

import (
	"fmt"
	"time"

	"obyoy-backend/contest/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storecontest "obyoy-backend/store/contest"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating contest
type Creater interface {
	Create(create *dto.Create) (*dto.CreateResponse, error)
}

// create creates contest
type create struct {
	storecontest storecontest.Contests
	validate     *validator.Validate
}

func (c *create) toModel(usercontest *dto.Create) (
	contest *model.Contest,
) {
	contest = &model.Contest{}
	contest.CreatedAt = time.Now().UTC()
	contest.UpdatedAt = contest.CreatedAt
	contest.ID = usercontest.ID
	contest.LandingURL = usercontest.LandingURL
	contest.ImageURL = usercontest.ImageURL
	contest.Name = usercontest.Name
	contest.Stadings = usercontest.Standings

	return
}

func (c *create) validateData(create *dto.Create) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Create) (
	modelcontest *model.Contest,
) {
	modelcontest = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Contest) (
	id string,
	err error,
) {
	id, err = c.storecontest.Save(model)
	return
}

func (c *create) giveResponse(modelcontest *model.Contest, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{}).Debug("User created contest successfully")

	return &dto.CreateResponse{
		Message: "contest created",
		OK:      true,
		//		contestTime: modelcontest.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create contest. Error: ", err)
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

	modelcontest := c.convertData(create)
	id, err := c.askStore(modelcontest)
	if err == nil {
		return c.giveResponse(modelcontest, id)
	}

	err = c.giveError()
	return nil, err
}

// CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	Storecontests storecontest.Contests
	Validate      *validator.Validate
}

// NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.Storecontests,
		params.Validate,
	}
}
