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
type Lister interface {
	List(req *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error)
}

// datasetReader implements Reader interface
type datasetLister struct {
	datasets storedataset.Datasets
}

func (list *datasetLister) askStore(state string, skip int64, limit int64) (
	dataset []*model.Dataset,
	err error,
) {
	dataset, err = list.datasets.FindByDatasetID(state, skip, limit)
	return
}

func (list *datasetLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *datasetLister) prepareResponse(
	datasets []*model.Dataset,
) (
	resp []dto.ReadResp,
) {
	for _, dataset := range datasets {
		var tmp dto.ReadResp
		tmp.FromModel(dataset)
		resp = append(resp, tmp)
	}
	//resp.FromModel(dataset)
	return
}

func (read *datasetLister) List(datasetReq *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	datasets, err := read.askStore(datasetReq.UserID, skip, limit)
	if err != nil {
		logrus.Error("Could not find dataset error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(datasets)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewListerParams struct {
	dig.In
	Dataset storedataset.Datasets
}

// NewReader provides Reader
func NewList(params NewReaderParams) Lister {
	return &datasetLister{
		datasets: params.Dataset,
	}
}
