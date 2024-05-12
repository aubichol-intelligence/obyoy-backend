package user

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storeuser "obyoy-backend/store/user"
	"obyoy-backend/user/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading users
type Lister interface {
	List(skip int64, limit int64, state string) ([]dto.ReadResp, error)
}

// userReader implements Reader interface
type restLister struct {
	users storeuser.Users
}

func (list *restLister) askStore(skip int64, limit int64, state string) (
	user []*model.User,
	err error,
) {
	user, err = list.users.List(skip, limit, state)
	return
}

func (list *restLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *restLister) prepareResponse(
	users []*model.User,
) (
	resp []dto.ReadResp,
) {
	for _, user := range users {
		var tmp dto.ReadResp
		tmp.FromModel(user)
		resp = append(resp, tmp)
	}
	return
}

func (read *restLister) List(skip int64, limit int64, state string) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	users, err := read.askStore(skip, limit, state)
	if err != nil {
		logrus.Error("Could not find user error : ", err)
		return nil, read.giveError()
	}
	var resp []dto.ReadResp
	resp = read.prepareResponse(users)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewStateListerParams struct {
	dig.In
	User storeuser.Users
}

// NewReader provides Reader
func NewLister(params NewStateListerParams) Lister {
	return &restLister{
		users: params.User,
	}
}
