package datastream

import (
	"obyoy-backend/datastream/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedatastream "obyoy-backend/store/datastream"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading datastreames
type Lister interface {
	List(req *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error)
}

// datastreamReader implements Reader interface
type datastreamLister struct {
	datastreams storedatastream.Datastreams
}

func (list *datastreamLister) askStore(state string, skip int64, limit int64) (
	datastream []*model.Datastream,
	err error,
) {
	datastream, err = list.datastreams.FindByDatastreamID(state, skip, limit)
	return
}

func (list *datastreamLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *datastreamLister) prepareResponse(
	datastreams []*model.Datastream,
) (
	resp []dto.ReadResp,
) {
	for _, datastream := range datastreams {
		var tmp dto.ReadResp
		tmp.FromModel(datastream)
		resp = append(resp, tmp)
	}
	//resp.FromModel(datastream)
	return
}

func (read *datastreamLister) List(datastreamReq *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	datastreams, err := read.askStore(datastreamReq.UserID, skip, limit)
	if err != nil {
		logrus.Error("Could not find datastream error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(datastreams)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewListerParams struct {
	dig.In
	Datastream storedatastream.Datastreams
}

// NewReader provides Reader
func NewList(params NewReaderParams) Lister {
	return &datastreamLister{
		datastreams: params.Datastream,
	}
}
