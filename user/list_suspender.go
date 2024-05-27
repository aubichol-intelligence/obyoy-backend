package user

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storeuser "obyoy-backend/store/user"
	"obyoy-backend/user/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading useres
type ListSuspender interface {
	ListSuspend(skip int64, limit int64) ([]dto.ReadResp, error)
}

// userReader implements Reader interface
type userLister struct {
	users storeuser.Users
}

func (list *userLister) askStore(skip int64, limit int64) (
	user []*model.User,
	err error,
) {
	user, err = list.users.ListSuspend(skip, limit)
	return
}

func (list *userLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *userLister) prepareResponse(
	users []*model.User,
) (
	resp []dto.ReadResp,
) {
	for _, user := range users {
		var tmp dto.ReadResp
		tmp.FromModel(user)
		resp = append(resp, tmp)
	}
	//resp.FromModel(user)
	return
}

func (read *userLister) ListSuspend(skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	users, err := read.askStore(skip, limit)
	if err != nil {
		logrus.Error("Could not find user error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(users)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewListerParams struct {
	dig.In
	User storeuser.Users
}

// NewReader provides Reader
func NewListSuspender(params NewListerParams) ListSuspender {
	return &userLister{
		users: params.User,
	}
}
