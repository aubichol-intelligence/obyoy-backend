package contest

import (
	"obyoy-backend/contest/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storecontest "obyoy-backend/store/contest"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading contestes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

// contestReader implements Reader interface
type contestReader struct {
	contests storecontest.Contests
}

func (read *contestReader) askStore(contestID string) (
	contest *model.Contest,
	err error,
) {
	contest, err = read.contests.FindByID(contestID)
	return
}

func (read *contestReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *contestReader) prepareResponse(
	contest *model.Contest,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(contest)
	return
}

func (read *contestReader) Read(contestReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	contest, err := read.askStore(contestReq.ContestID)
	if err != nil {
		logrus.Error("Could not find contest error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(contest)

	return &resp, nil
}

// NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Contest storecontest.Contests
}

// NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &contestReader{
		contests: params.Contest,
	}
}
