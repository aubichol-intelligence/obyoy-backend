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
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

// parallelsentenceReader implements Reader interface
type parallelsentenceReader struct {
	parallelsentences storeparallelsentence.Parallelsentences
}

func (read *parallelsentenceReader) askStore(parallelsentenceID string) (
	parallelsentence *model.Parallelsentence,
	err error,
) {
	parallelsentence, err = read.parallelsentences.FindByID(parallelsentenceID)
	return
}

func (read *parallelsentenceReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *parallelsentenceReader) prepareResponse(
	parallelsentence *model.Parallelsentence,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(parallelsentence)
	return
}

func (read *parallelsentenceReader) Read(parallelsentenceReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	parallelsentence, err := read.askStore(parallelsentenceReq.ParallelsentenceID)
	if err != nil {
		logrus.Error("Could not find parallelsentence error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(parallelsentence)

	return &resp, nil
}

// NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Parallelsentence storeparallelsentence.Parallelsentences
}

// NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &parallelsentenceReader{
		parallelsentences: params.Parallelsentence,
	}
}
