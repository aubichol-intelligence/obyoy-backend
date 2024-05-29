package parallelsentence

import (
	"obyoy-backend/errors"
	"obyoy-backend/parallelsentence/dto"
	"obyoy-backend/store/parallelsentence"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// CountByStatusReader provides an interface for counting
// parallelsentences of a status
type CountReader interface {
	Count(*dto.CountReq) (*dto.CountResp, error)
}

// countByUserReader implements Reader interface
type countByStatusReader struct {
	parallelsentences parallelsentence.Parallelsentences
}

func (count *countByStatusReader) askStore(StatusID string) (
	counts int64,
	err error,
) {
	counts, err = count.parallelsentences.Count()
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
		message := "Could not count parallelsentence by status id error : "
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
	Parallelsentence parallelsentence.Parallelsentences
}

// NewCountByStatusReader provides CountByStatusReader
func NewCountByStatusReader(
	params NewCountParams,
) CountReader {
	return &countByStatusReader{
		parallelsentences: params.Parallelsentence,
	}
}
