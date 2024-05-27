package dataset

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/dataset"
	"obyoy-backend/dataset/dto"
	"obyoy-backend/errors"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// listHandler holds the handler that searches for dataset
type listHandler struct {
	searcher dataset.Lister
}

func (ch *listHandler) decodeBody(
	body io.ReadCloser,
) (
	dataset dto.ListReq,
	err error,
) {
	err = dataset.FromReader(body)
	return
}

// querySkip skips number of datasets for a request
func (list *listHandler) querySkip(
	r *http.Request,
) (skip int, err error) {
	skip, err = strconv.Atoi(chi.URLParam(r, "skip"))
	if err != nil {
		err = fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{
					"Invalid skip", false,
				},
			},
		)
		return
	}
	return
}

// queryLimit limits number of datasets per query
func (list *listHandler) queryLimit(r *http.Request) (
	limit int, err error,
) {
	limit, err = strconv.Atoi(chi.URLParam(r, "limit"))

	if err != nil {
		err = fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{
					"Invalid limit", false,
				},
			},
		)
		return
	}

	if limit > 50 {
		err = &errors.Invalid{
			errors.Base{
				"Limit is too big", false,
			},
		}
		return
	}

	return
}

func (list *listHandler) askController(
	listReq dto.ListReq,
) (
	resp []dto.ReadResp,
	err error,
) {
	resp, err = list.searcher.List(&listReq, listReq.Skip, listReq.Limit)
	return
}

func (list *listHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (list *listHandler) responseSuccess(
	w http.ResponseWriter,
	resp []dto.ReadResp,
) {
	// Serve a response to the client
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

func (list *listHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	var listDat dto.ListReq

	skip, err := list.querySkip(r)
	limit, err := list.queryLimit(r)

	listDat.Skip = int64(skip)
	listDat.Limit = int64(limit)

	// Read request from database using skip and limit
	resp, err := list.askController(listDat)

	if err != nil {
		list.handleError(w, err)
		return
	}

	list.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler
func (list *listHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	list.handleRead(w, r)
}

// ListRoute provides a route that gives lists for datasets
func ListRoute(
	lister dataset.Lister,
	middleware *middleware.Auth,
) *routeutils.Route {
	handler := listHandler{lister}
	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.DatesetList,
		Handler: middleware.Middleware(&handler),
	}
}
