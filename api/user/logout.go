package user

import (
	"io"
	"net/http"

	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/user"
	"obyoy-backend/user/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// loginHandler holds login handler
type logoutHandler struct {
	userLogin user.Login
}

func (lh *logoutHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (lh *logoutHandler) responseSuccess(
	w http.ResponseWriter,
) {
	var resp dto.Token
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

func (lh *logoutHandler) decodeBody(body io.ReadCloser) (
	login dto.Login, err error,
) {
	err = login.FromReader(body)
	return
}

// ServeHTTP implements http.Handler interface for loginHandler
func (lh *logoutHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	lh.responseSuccess(w)
}

// LoginRouteParams lists all the paramters for NewLoginRoute
type LogoutRouteParams struct {
	dig.In
	UserLogin user.Login
}

// LoginRoute provides a route that gives a user login options
func LogoutRoute(
	params LoginRouteParams,
) *routeutils.Route {
	handler := loginHandler{params.UserLogin}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.LogoutUser,
		Handler: &handler,
	}
}
