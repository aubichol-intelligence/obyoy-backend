package datastream

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/datastream"
	"obyoy-backend/datastream/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// createHandler holds handler for creating datastream items
type createHandler struct {
	create datastream.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	datastream dto.Create,
	err error,
) {
	err = datastream.FromReader(body)
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
	datastream *dto.Create,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(datastream)
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

// ServeHTTP implements http.Handler interface for datastream create
func (ch *createHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	datastreamDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	//	datastreamDat.UserID = ch.decodeContext(r)
	data, err := ch.askController(&datastreamDat)

	if err != nil {
		message := "Unable to create datastream error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     datastream.Creater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to create datastreams
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DatastreamCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
