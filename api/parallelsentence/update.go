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

type updateHandler struct {
	reader parallelsentence.Reader
}

func (read *updateHandler) decodeURL(
	r *http.Request,
) (parallelsentenceID string) {
	// Get user id from url
	parallelsentenceID = chi.URLParam(r, "id")
	return
}

func (read *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (read *updateHandler) askController(
	req *dto.ReadReq,
) (
	resp *dto.ReadResp,
	err error,
) {
	resp, err = read.reader.Read(req)
	return
}

func (read *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (read *updateHandler) responseSuccess(
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

func (read *updateHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.ReadReq{}
	//	req.parallelsentenceID = read.decodeURL(r)

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
func (read *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	read.handleRead(w, r)
}

// ReadRouteParams lists all the parameters for ReadRoute
type UpdateRouteParams struct {
	dig.In
	Reader     parallelsentence.Reader
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get a parallelsentence item
func UpdateRoute(params ReadRouteParams) *routeutils.Route {

	handler := updateHandler{
		reader: params.Reader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.ParallelsentenceUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
