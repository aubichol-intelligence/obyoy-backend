package user

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/errors"
	"obyoy-backend/user"
	"obyoy-backend/user/dto"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// searchHandler holds the handler that searches for restaurant
type listHandler struct {
	searcher user.Lister
}

func (ch *listHandler) decodeBody(
	body io.ReadCloser,
) (
	restaurant dto.ListByType,
	err error,
) {
	err = restaurant.FromReader(body)
	return
}

// querySkip skips number of users for a request
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

// queryLimit limits number of users per query
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
	skip int64,
	limit int64,
	state string,
) (
	resp []dto.ReadResp,
	err error,
) {
	resp, err = list.searcher.List(skip, limit, state)
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

	listDat, err := list.decodeBody(r.Body)

	if err != nil {
		//message := "Unable to decode error: "
		list.handleError(w, err)
		return
	}

	// Read request from database using request id and user id
	resp, err := list.askController(listDat.Skip, listDat.Limit, listDat.State)

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

// SearchRoute provides a route that searches for restaurants
func ListRoute(
	searcher user.Lister,
	middleware *middleware.Auth,
) *routeutils.Route {
	handler := listHandler{searcher}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.UserList,
		Handler: middleware.Middleware(&handler),
	}
}
