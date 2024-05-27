package translation

import (
	"fmt"
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/translation"
	"obyoy-backend/translation/dto"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type updateHandler struct {
	updater translation.Updater
}

func (ch *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	translation dto.Update,
	err error,
) {
	err = translation.FromReader(body)
	return
}

func (update *updateHandler) decodeURL(
	r *http.Request,
) (translationID string) {
	// Get user id from url
	translationID = chi.URLParam(r, "id")
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

	req, err := update.decodeBody(r.Body)

	fmt.Println(err)
	//req := dto.Update{}
	//	req.translationID = read.decodeURL(r)

	//	req.UserID = update.decodeContext(r)

	// Read request from database using request id and user id
	resp, err := update.askController(&req)

	if err != nil {
		update.handleError(w, err)
		return
	}

	update.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler for updating translation
func (update *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	update.handleRead(w, r)
}

// UpdateRouteParams lists all the parameters for UpdateRoute
type UpdateRouteParams struct {
	dig.In
	Updater    translation.Updater
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get a translation item
func UpdateRoute(params UpdateRouteParams) *routeutils.Route {

	handler := updateHandler{
		updater: params.Updater,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.TranslationUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
