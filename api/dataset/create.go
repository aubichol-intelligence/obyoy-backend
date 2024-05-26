package dataset

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/dataset"
	"obyoy-backend/dataset/dto"
	"obyoy-backend/datastream"
	datastreamdto "obyoy-backend/datastream/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// createHandler holds handler for creating dataset items
type createHandler struct {
	create       dataset.Creater
	createstream datastream.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	dataset dto.Create,
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
	dataset *dto.Create,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(dataset)

	for ind, line := range dataset.Set {
		var one_line datastreamdto.Create
		one_line.DatasetID = data.ID
		one_line.LineNumber = int32(ind + 1)
		one_line.SourceSentence = line
		one_line.Name = dataset.Name
		ch.createstream.Create(&one_line)
	}

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

	datasetDat.UploaderID = ch.decodeContext(r)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

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
	Create           dataset.Creater
	CreateDataStream datastream.Creater
	Middleware       *middleware.Auth
}

// CreateRoute provides a route that lets to take datasets
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create, params.CreateDataStream}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DatasetCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
