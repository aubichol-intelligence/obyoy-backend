package dataset

import (
	"obyoy-backend/dataset/dto"
	"obyoy-backend/errors"
	"obyoy-backend/store/dataset"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// CountByStatusReader provides an interface for counting
// datasets of a status
type CountReader interface {
	Count(*dto.CountReq) (*dto.CountResp, error)
}

// countByUserReader implements Reader interface
type countByStatusReader struct {
	datasets dataset.Datasets
}

func (count *countByStatusReader) askStore(StatusID string) (
	counts int64,
	err error,
) {
	counts, err = count.datasets.Count()
	return
}

func (count *countByStatusReader) logError(
	message string,
	err error,
) {
	logrus.Error(message, err)
}

func (count *countByStatusReader) giveError() (err error) {
	return &errors.Unknown{
		errors.Base{"Invalid request", false},
	}
}

func (count *countByStatusReader) prepareResopnse(counts int64) (
	resp dto.CountResp,
) {
	resp.FromModel(counts)
	return
}

func (count *countByStatusReader) giveResponse(
	counts dto.CountResp,
) (
	*dto.CountResp,
	error,
) {
	return &counts, nil
}

func (count *countByStatusReader) giveErrorResponse(err error) (
	*dto.CountResp,
	error,
) {
	return nil, err
}

// Count implements CountByStatusReader interface
func (count *countByStatusReader) Count(
	countByStatusReq *dto.CountReq,
) (*dto.CountResp, error) {

	counts, err := count.askStore(
		countByStatusReq.StatusID,
	)
	if err != nil {
		message := "Could not count dataset by status id error : "
		count.logError(message, err)
		err = count.giveError()
		return count.giveErrorResponse(err)
	}

	var resp dto.CountResp
	resp = count.prepareResopnse(counts)
	return count.giveResponse(resp)
}

// NewCountByStatusReaderParams lists params for the
// NewCountByStatusReader
type NewCountParams struct {
	dig.In
	Dataset dataset.Datasets
}

// NewCountByStatusReader provides CountByStatusReader
func NewCountByStatusReader(
	params NewCountParams,
) CountReader {
	return &countByStatusReader{
		datasets: params.Dataset,
	}
}
