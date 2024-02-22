package menuitem

import (
	"net/http"

	"horkora-backend/api/middleware"
	"horkora-backend/api/routeutils"
	"horkora-backend/apipattern"
	"horkora-backend/menuitem"
	"horkora-backend/menuitem/dto"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type readHandler struct {
	reader menuitem.Reader
}

func (read *readHandler) decodeURL(
	r *http.Request,
) (menuitemID string) {
	// Get user id from url
	menuitemID = chi.URLParam(r, "id")
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
	req.MenuItemID = read.decodeURL(r)

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
	Reader     menuitem.Reader
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get a menu item
func ReadRoute(params ReadRouteParams) *routeutils.Route {

	handler := readHandler{
		reader: params.Reader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.MenuItemRead,
		Handler: params.Middleware.Middleware(&handler),
	}
}
(base) nelson@NELSONs-MacBook-Pro menuitem % cat update.go
package menuitem

import (
	"io"
	"net/http"

	"horkora-backend/api/middleware"
	"horkora-backend/api/routeutils"
	"horkora-backend/apipattern"
	"horkora-backend/menuitem"
	"horkora-backend/menuitem/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// updateHandler holds menu item update handler
type updateHandler struct {
	update menuitem.Updater
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

	menuitemDat := dto.Update{}
	menuitemDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode menuitem update error: "
		ch.handleError(w, err, message)
		return
	}

	menuitemDat.UserID = ch.decodeContext(r)

	data, err := ch.askController(&menuitemDat)

	if err != nil {
		message := "Unable to update menuitem error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// UpdateParams provide parameters for menuitem update handler
type UpdateParams struct {
	dig.In
	Update     menuitem.Updater
	Middleware *middleware.Auth
}

// UpdateRoute provides a route that updates a menuitem
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.MenuItemUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
