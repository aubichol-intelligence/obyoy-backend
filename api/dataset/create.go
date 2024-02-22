package dataset

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/dataset"
	"obyoy-backend/dataset/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// createHandler holds handler for creating dataset items
type createHandler struct {
	create dataset.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	dataset dto.dataset,
	err error,
) {
	err = dataset.FromReader(body)
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
	dataset *dto.dataset,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(dataset)
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

	datasetDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	datasetDat.UserID = ch.decodeContext(r)
	data, err := ch.askController(&datasetDat)

	if err != nil {
		message := "Unable to create dataset error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     dataset.Creater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to take datasets
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.datasetCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}