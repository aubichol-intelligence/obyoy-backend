package dataset

import (
	"obyoy-backend/dataset/dto"
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storedataset "obyoy-backend/store/dataset"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading datasetes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

// datasetReader implements Reader interface
type datasetReader struct {
	datasets storedataset.Datasets
}

func (read *datasetReader) askStore(datasetID string) (
	dataset *model.Dataset,
	err error,
) {
	dataset, err = read.datasets.FindByID(datasetID)
	return
}

func (read *datasetReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *datasetReader) prepareResponse(
	dataset *model.Dataset,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(dataset)
	return
}

func (read *datasetReader) Read(datasetReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	dataset, err := read.askStore(datasetReq.DatasetID)
	if err != nil {
		logrus.Error("Could not find dataset error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(dataset)

	return &resp, nil
}

// NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Dataset storedataset.Datasets
}

// NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &datasetReader{
		datasets: params.Dataset,
	}
}
