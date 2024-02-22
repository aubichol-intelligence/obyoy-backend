package user

import (
	"net/http"

	"horkora-backend/api/middleware"
	"horkora-backend/api/routeutils"
	"horkora-backend/apipattern"
	"horkora-backend/errors"
	"horkora-backend/user"
	"horkora-backend/user/dto"

	"github.com/sirupsen/logrus"
)

// searchHandler holds the handler that searches for user
type searchHandler struct {
	searcher user.Searcher
}

func (sh *searchHandler) query(r *http.Request) (
	q string, err error,
) {
	q = r.URL.Query().Get("q")
	if q == "" {
		err = &errors.Invalid{
			errors.Base{
				"Invalid search query", false,
			},
		}
	}

	return
}

func (sh *searchHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (sh *searchHandler) resopnseSuccess(
	w http.ResponseWriter,
	resp []*dto.Search,
) {
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

// ServeHTTP implements the http.Handler interface for
// user search handler
func (sh *searchHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()
	q, err := sh.query(r)
	if err != nil {
		message := ""
		sh.handleError(w, err, message)
		return
	}

	data, err := sh.searcher.Search(q)
	if err != nil {
		message := "Unable to search user, err: "
		sh.handleError(w, err, message)
		return
	}

	sh.resopnseSuccess(w, data)
}

// SearchRoute provides a route that searches for users
func SearchRoute(
	searcher user.Searcher,
	middleware *middleware.Auth,
) *routeutils.Route {
	handler := searchHandler{searcher}
	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.UserSearch,
		Handler: middleware.Middleware(&handler),
	}
}
