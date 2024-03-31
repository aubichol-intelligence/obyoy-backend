package datastream

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/datastream"
	"obyoy-backend/datastream/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// deleteHandler holds datastream item update handler
type deleteHandler struct {
	delete datastream.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	datastream dto.Delete,
	err error,
) {
	err = datastream.FromReader(body)
	return
}

func (dh *deleteHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (dh *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (dh *deleteHandler) askController(update *dto.Delete) (
	resp *dto.DeleteResponse,
	err error,
) {
	resp, err = dh.delete.Delete(update)
	return
}

func (dh *deleteHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.DeleteResponse,
) {
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

// ServeHTTP implements http.Handler interface
func (dh *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	datastream := dto.Delete{}
	datastream, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode datastream item delete error: "
		dh.handleError(w, err, message)
		return
	}

	datastream.UserID = dh.decodeContext(r)

	data, err := dh.askController(&datastream)

	if err != nil {
		message := "Unable to delete datastream item error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

// DeleteParams provide parameters for datastream delete handler
type DeleteParams struct {
	dig.In
	Delete     datastream.Deleter
	Middleware *middleware.Auth
}

// DeleteRoute provides a route that deletes datastream
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.datastreamDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
