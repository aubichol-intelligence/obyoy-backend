package parallelsentence

import (
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

type readHandler struct {
	reader parallelsentence.Reader
}

func (read *readHandler) decodeURL(
	r *http.Request,
) (parallelsentenceID string) {
	// Get user id from url
	parallelsentenceID = chi.URLParam(r, "id")
	return
}

func (read *readHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (read *readHandler) askController(
	req *dto.ReadReq,
) (
	resp *dto.ReadResp,
	err error,
) {
	resp, err = read.reader.Read(req)
	return
}

func (read *readHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (read *readHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.ReadResp,
) {
	// Serve a response to the client
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

func (read *readHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.ReadReq{}
	req.parallelsentenceID = read.decodeURL(r)

	req.UserID = read.decodeContext(r)

	// Read request from database using request id and user id
	resp, err := read.askController(&req)

	if err != nil {
		read.handleError(w, err)
		return
	}

	read.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler
func (read *readHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	read.handleRead(w, r)
}

// ReadRouteParams lists all the parameters for ReadRoute
type ReadRouteParams struct {
	dig.In
	Reader     parallelsentence.Reader
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get a parallelsentence item
func ReadRoute(params ReadRouteParams) *routeutils.Route {

	handler := readHandler{
		reader: params.Reader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.parallelsentenceRead,
		Handler: params.Middleware.Middleware(&handler),
	}
}
(base) nelson@NELSONs-MacBook-Pro parallelsentence % cat update.go
package parallelsentence

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/parallelsentence"
	"obyoy-backend/parallelsentence/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// updateHandler holds parallelsentence item update handler
type updateHandler struct {
	update parallelsentence.Updater
}

func (uh *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	bloodreqAtt dto.Update,
	err error,
) {
	err = bloodreqAtt.FromReader(body)
	return
}

func (uh *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (uh *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (uh *updateHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = uh.update.Update(update)
	return
}

func (uh *updateHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.UpdateResponse,
) {
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

// ServeHTTP implements http.Handler interface
func (ch *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	parallelsentenceDat := dto.Update{}
	parallelsentenceDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode parallelsentence update error: "
		ch.handleError(w, err, message)
		return
	}

	parallelsentenceDat.UserID = ch.decodeContext(r)

	data, err := ch.askController(&parallelsentenceDat)

	if err != nil {
		message := "Unable to update parallelsentence error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// UpdateParams provide parameters for parallelsentence update handler
type UpdateParams struct {
	dig.In
	Update     parallelsentence.Updater
	Middleware *middleware.Auth
}

// UpdateRoute provides a route that updates a parallelsentence
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.parallelsentenceUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
