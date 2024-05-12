package contest

import (
	"ardent-backend/contest/dto"
	"ardent-backend/errors"
	"ardent-backend/model"
	storecontest "ardent-backend/store/contest"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading contestes
type Lister interface {
	List(req *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error)
}

// contestReader implements Reader interface
type contestLister struct {
	contests storecontest.Contests
}

func (list *contestLister) askStore(state string, skip int64, limit int64) (
	contest []*model.Contest,
	err error,
) {
	contest, err = list.contests.FindByContestID(state, skip, limit)
	return
}

func (list *contestLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *contestLister) prepareResponse(
	contests []*model.Contest,
) (
	resp []dto.ReadResp,
) {
	for _, contest := range contests {
		var tmp dto.ReadResp
		tmp.FromModel(contest)
		resp = append(resp, tmp)
	}
	//resp.FromModel(contest)
	return
}

func (read *contestLister) List(contestReq *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	contests, err := read.askStore(contestReq.UserID, skip, limit)
	if err != nil {
		logrus.Error("Could not find contest error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(contests)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewListerParams struct {
	dig.In
	Contest storecontest.Contests
}

// NewReader provides Reader
func NewList(params NewReaderParams) Lister {
	return &contestLister{
		contests: params.Contest,
	}
}
