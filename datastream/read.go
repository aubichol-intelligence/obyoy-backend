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
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

// datastreamReader implements Reader interface
type datastreamReader struct {
	datastreams storedatastream.Datastreams
}

func (read *datastreamReader) askStore(datastreamID string) (
	datastream *model.Datastream,
	err error,
) {
	datastream, err = read.datastreams.FindByID(datastreamID)
	return
}

func (read *datastreamReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *datastreamReader) prepareResponse(
	datastream *model.Datastream,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(datastream)
	return
}

func (read *datastreamReader) Read(datastreamReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	datastream, err := read.askStore(datastreamReq.DatastreamID)
	if err != nil {
		logrus.Error("Could not find datastream error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(datastream)

	return &resp, nil
}

// NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Datastream storedatastream.Datastreams
}

// NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &datastreamReader{
		datastreams: params.Datastream,
	}
}
