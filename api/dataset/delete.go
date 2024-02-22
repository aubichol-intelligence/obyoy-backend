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

// deleteHandler holds menu item update handler
type deleteHandler struct {
	delete menuitem.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	menuitem dto.Delete,
	err error,
) {
	err = menuitem.FromReader(body)
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

	menuitem := dto.Delete{}
	menuitem, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode menu item delete error: "
		dh.handleError(w, err, message)
		return
	}

	menuitem.UserID = dh.decodeContext(r)

	data, err := dh.askController(&menuitem)

	if err != nil {
		message := "Unable to delete menu item error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

// DeleteParams provide parameters for menuitem delete handler
type DeleteParams struct {
	dig.In
	Delete     menuitem.Deleter
	Middleware *middleware.Auth
}

// DeleteRoute provides a route that deletes menuitem
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.MenuItemDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
