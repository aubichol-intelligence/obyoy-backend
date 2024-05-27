package parallelsentence

import (
	"fmt"
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/parallelsentence"
	"obyoy-backend/parallelsentence/dto"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type updateHandler struct {
	updater parallelsentence.Updater
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
) (parallelsentenceID string) {
	// Get user id from url
	parallelsentenceID = chi.URLParam(r, "id")
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

func (update *updateHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.Update{}
	var err error
	req, err = update.decodeBody(r.Body)

	fmt.Println(err)

	//	req.UserID = update.decodeContext(r)

	// Read request from database using request id and user id
	resp, err := update.askController(&req)

	if err != nil {
		update.handleError(w, err)
		return
	}

	update.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler for updating a parallel sentence
func (update *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	update.handleRead(w, r)
}

// UpdateRouteParams lists all the parameters for ReadRoute
type UpdateRouteParams struct {
	dig.In
	Updater    parallelsentence.Updater
	Middleware *middleware.Auth
}

// UpdateRoute provides a route to get a parallelsentence item
func UpdateRoute(params UpdateRouteParams) *routeutils.Route {

	handler := updateHandler{
		updater: params.Updater,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.ParallelsentenceUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
