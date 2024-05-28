package parallelsentence

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	"obyoy-backend/parallelsentence/dto"
	storeparallelsentence "obyoy-backend/store/parallelsentence"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading parallelsentencees
type Lister interface {
	List(req *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error)
}

// parallelsentenceReader implements Reader interface
type parallelsentenceLister struct {
	parallelsentences storeparallelsentence.Parallelsentences
}

func (list *parallelsentenceLister) askStore(state string, skip int64, limit int64) (
	parallelsentence []*model.Parallelsentence,
	err error,
) {
	parallelsentence, err = list.parallelsentences.List(skip, limit)
	return
}

func (list *parallelsentenceLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *parallelsentenceLister) prepareResponse(
	parallelsentences []*model.Parallelsentence,
) (
	resp []dto.ReadResp,
) {
	for _, parallelsentence := range parallelsentences {
		var tmp dto.ReadResp
		tmp.FromModel(parallelsentence)
		resp = append(resp, tmp)
	}
	//resp.FromModel(parallelsentence)
	return
}

func (read *parallelsentenceLister) List(parallelsentenceReq *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	parallelsentences, err := read.askStore(parallelsentenceReq.UserID, skip, limit)
	if err != nil {
		logrus.Error("Could not find parallelsentence error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(parallelsentences)

	return resp, nil
}

// NewListerParams lists params for the NewList
type NewListerParams struct {
	dig.In
	Parallelsentence storeparallelsentence.Parallelsentences
}

// NewList provides Reader
func NewList(params NewListerParams) Lister {
	return &parallelsentenceLister{
		parallelsentences: params.Parallelsentence,
	}
}
