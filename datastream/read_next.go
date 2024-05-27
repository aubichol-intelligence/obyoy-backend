package datastream

import (
	"obyoy-backend/datastream/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedatastream "obyoy-backend/store/datastream"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading the next line from the datastreames
type NextReader interface {
	ReadNext(*dto.ReadReq) (*dto.ReadResp, error)
}

// datastreamReader implements Reader interface
type datastreamNextReader struct {
	datastreams storedatastream.Datastreams
}

func (read *datastreamNextReader) askStore() (
	datastream *model.Datastream,
	err error,
) {
	datastream, err = read.datastreams.FindNext()
	return
}

func (read *datastreamNextReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *datastreamNextReader) prepareResponse(
	datastream *model.Datastream,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(datastream)
	return
}

func (read *datastreamNextReader) ReadNext(datastreamReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	datastream, err := read.askStore()

	if err != nil {
		logrus.Error("Could not find datastream error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(datastream)

	return &resp, nil
}

// NewReaderNextParams lists params for the NewReader
type NewReaderNextParams struct {
	dig.In
	Datastream storedatastream.Datastreams
}

// NewNextReader provides NextReader
func NewNextReader(params NewReaderNextParams) NextReader {
	return &datastreamNextReader{
		datastreams: params.Datastream,
	}
}
