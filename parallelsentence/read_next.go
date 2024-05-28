package parallelsentence

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	"obyoy-backend/parallelsentence/dto"
	storeparallelsentence "obyoy-backend/store/parallelsentence"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading the next line from the parallelsentencees
type NextReader interface {
	ReadNext(*dto.ReadReq) (*dto.ReadResp, error)
}

// parallelsentenceReader implements Reader interface
type parallelsentenceNextReader struct {
	parallelsentences storeparallelsentence.Parallelsentences
}

func (read *parallelsentenceNextReader) askStore() (
	parallelsentence *model.Parallelsentence,
	err error,
) {
	parallelsentence, err = read.parallelsentences.FindNext()
	return
}

func (read *parallelsentenceNextReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *parallelsentenceNextReader) prepareResponse(
	parallelsentence *model.Parallelsentence,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(parallelsentence)
	return
}

func (read *parallelsentenceNextReader) ReadNext(parallelsentenceReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	parallelsentence, err := read.askStore()

	if err != nil {
		logrus.Error("Could not find parallelsentence error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(parallelsentence)

	return &resp, nil
}

// NewReaderNextParams lists params for the NewReader
type NewReaderNextParams struct {
	dig.In
	Parallelsentence storeparallelsentence.Parallelsentences
}

// NewNextReader provides NextReader
func NewNextReader(params NewReaderNextParams) NextReader {
	return &parallelsentenceNextReader{
		parallelsentences: params.Parallelsentence,
	}
}
