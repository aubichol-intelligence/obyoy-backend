package parallelsentence

import (
	"fmt"
	"net/http"
	"strconv"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"

	"obyoy-backend/apipattern"

	"obyoy-backend/parallelsentence"
	"obyoy-backend/parallelsentence/dto"

	"obyoy-backend/errors"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type countHandler struct {
	countByStatusReader parallelsentence.CountReader
}

// querySkip skips number of parallelsentences for a request
func (read *countHandler) querySkip(
	r *http.Request,
) (
	skip int,
	err error,
) {
	skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	if err != nil {
		err = fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{errors.Base{"Invalid skip", false}},
		)
		return
	}
	return
}

// queryLimit limits number of parallelsentences per query
func (read *countHandler) queryLimit(
	r *http.Request,
) (
	limit int,
	err error,
) {
	limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		err = fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{errors.Base{
				"Invalid limit",
				false,
			},
			},
		)
		return
	}

	if limit < 0 {
		err = &errors.Invalid{
			errors.Base{
				"Limit is too small",
				false,
			},
		}
		return
	}

	if limit > 50 {
		err = &errors.Invalid{
			errors.Base{
				"Limit is too big",
				false,
			},
		}
		return
	}

	return
}

func (read *countHandler) decodeURL(
	r *http.Request,
) (statusID string) {
	statusID = chi.URLParam(r, "id")
	return
}

func (read *countHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (read *countHandler) askController(
	message *dto.CountReq,
) (resp *dto.CountResp, err error) {
	resp, err = read.countByStatusReader.Count(message)
	return
}

func (read *countHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.CountResp,
) {
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

func (read *countHandler) handleCountParallelsentence(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.CountReq{}

	req.StatusID = read.decodeURL(r)
	// Read status from database using status id and user id
	// TO-DO: skip and limit may not be needed here.

	resp, err := read.askController(&req)

	if err != nil {
		read.handleError(w, err)
		return
	}

	// Serve a response to the client
	read.responseSuccess(w, resp)

}

// ServeHTTP implements http.Handler interface of a parallelsentence
func (read *countHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	read.handleCountParallelsentence(w, r)
}

/*
CountByStatusParams lists all the parameters
for CountByStatusRoute
*/
type CountParams struct {
	dig.In
	CountByStatus parallelsentence.CountReader
	Middleware    *middleware.Auth
}

/*
CountByStatusRoute provides a route to count
messages from status
*/
func CountRoute(
	params CountParams,
) *routeutils.Route {

	handler := countHandler{
		countByStatusReader: params.CountByStatus,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.ParallelsentenceCount,
		Handler: params.Middleware.Middleware(&handler),
	}
}
