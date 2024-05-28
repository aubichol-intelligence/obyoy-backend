package parallelsentence

import (
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/parallelsentence"
	"obyoy-backend/parallelsentence/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type readNextHandler struct {
	reader parallelsentence.NextReader
}

func (read *readNextHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (read *readNextHandler) askController(
	req *dto.ReadReq,
) (
	resp *dto.ReadResp,
	err error,
) {
	resp, err = read.reader.ReadNext(req)
	return
}

func (read *readNextHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (read *readNextHandler) responseSuccess(
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

func (read *readNextHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.ReadReq{}

	req.UserID = read.decodeContext(r)

	// Read request from database using request id and user id
	resp, err := read.askController(&req)

	if err != nil {
		read.handleError(w, err)
		return
	}

	read.responseSuccess(w, resp)
}

// ServeHTTP implements http.Handler for reading the next line from the parallelsentence
func (read *readNextHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	read.handleRead(w, r)
}

// ReadRouteParams lists all the parameters for ReadRoute
type ReadNextRouteParams struct {
	dig.In
	NextReader parallelsentence.NextReader
	Middleware *middleware.Auth
}

// ReadRoute provides a route to get the next parallelsentence item
func ReadNextRoute(params ReadNextRouteParams) *routeutils.Route {

	handler := readNextHandler{
		reader: params.NextReader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.ParallelsentenceReadNext,
		Handler: params.Middleware.Middleware(&handler),
	}
}
