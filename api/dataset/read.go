package dataset

import (
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

type readHandler struct {
	reader dataset.Reader
}

func (read *readHandler) decodeURL(
	r *http.Request,
) (datasetID string) {
	// Get dataset id from url
	datasetID = chi.URLParam(r, "id")
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
	// Serve a successful response to the client
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

	req.DatasetID = read.decodeURL(r)
	req.UserID = read.decodeContext(r)

	// Read request from database using dataset id and user id
	resp, err := read.askController(&req)

	if err != nil {
		read.handleError(w, err)
		return
	}

	read.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler for dataset read
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
	Reader     dataset.Reader
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get a dataset item
func ReadRoute(params ReadRouteParams) *routeutils.Route {

	handler := readHandler{
		reader: params.Reader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.DatasetRead,
		Handler: params.Middleware.Middleware(&handler),
	}
}
