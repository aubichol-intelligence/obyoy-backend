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

// createHandler holds handler for creating menu items
type createHandler struct {
	create menuitem.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	menuitem dto.MenuItem,
	err error,
) {
	err = menuitem.FromReader(body)
	return
}

func (ch *createHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (ch *createHandler) askController(
	menuitem *dto.MenuItem,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(menuitem)
	return
}

func (ch *createHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *createHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.CreateResponse,
) {
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

// ServeHTTP implements http.Handler interface
func (ch *createHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	menuitemDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	menuitemDat.UserID = ch.decodeContext(r)
	data, err := ch.askController(&menuitemDat)

	if err != nil {
		message := "Unable to create menuitem error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     menuitem.Creater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to take menuitems
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.MenuItemCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}