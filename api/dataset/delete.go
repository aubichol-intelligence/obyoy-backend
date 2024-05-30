package dataset

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/dataset"
	"obyoy-backend/dataset/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// deleteHandler holds dataset item delete handler
type deleteHandler struct {
	delete dataset.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	dataset dto.Delete,
	err error,
) {
	err = dataset.FromReader(body)
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

	dataset := dto.Delete{}
	dataset, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode dataset item delete error: "
		dh.handleError(w, err, message)
		return
	}

	dataset.UserID = dh.decodeContext(r)

	data, err := dh.askController(&dataset)

	if err != nil {
		message := "Unable to delete dataset item error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

// DeleteParams provide parameters for dataset delete handler
type DeleteParams struct {
	dig.In
	Delete     dataset.Deleter
	Middleware *middleware.AdminAuth
}

// DeleteRoute provides a route that deletes dataset
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DatasetDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
