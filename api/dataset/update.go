package dataset

import (
	"fmt"
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/dataset"
	"obyoy-backend/dataset/dto"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type updateHandler struct {
	updater dataset.Updater
}

func (ch *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	dataset dto.Update,
	err error,
) {
	err = dataset.FromReader(body)
	return
}

func (update *updateHandler) decodeURL(
	r *http.Request,
) (datasetID string) {
	// Get user id from url
	datasetID = chi.URLParam(r, "id")
	return
}

func (update *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (update *updateHandler) askController(
	req *dto.Update,
) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = update.updater.Update(req)
	return
}

func (update *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (update *updateHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.UpdateResponse,
) {
	// Serve a response to the client
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

func (update *updateHandler) handleUpdate(
	w http.ResponseWriter,
	r *http.Request,
) {

	defer r.Body.Close()

	req, err := update.decodeBody(r.Body)

	fmt.Println(err)

	//	req.UserID = update.decodeContext(r)

	// Read request from database using dataset id and user id
	resp, err := update.askController(&req)

	if err != nil {
		update.handleError(w, err)
		return
	}

	update.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler for dataset update
func (update *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	update.handleUpdate(w, r)
}

// UpdateRouteParams lists all the parameters for UpdateRoute
type UpdateRouteParams struct {
	dig.In
	Updater    dataset.Updater
	Middleware *middleware.Auth
}

// UpdateRoute provides a route to get a dataset item
func UpdateRoute(params UpdateRouteParams) *routeutils.Route {

	handler := updateHandler{
		updater: params.Updater,
	}

	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DatasetUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
