package parallelsentence

import (
	"fmt"
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/datastream"
	datastreamDto "obyoy-backend/datastream/dto"
	"obyoy-backend/parallelsentence"
	"obyoy-backend/parallelsentence/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// createHandler holds handler for creating parallelsentence items
type createHandler struct {
	create parallelsentence.Creater
	update datastream.Updater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	parallelsentence dto.Create,
	err error,
) {
	err = parallelsentence.FromReader(body)
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
	parallelsentence *dto.Create,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(parallelsentence)
	fmt.Println(err)
	var datastream datastreamDto.Update
	datastream.ID = parallelsentence.DatastreamID
	datastream.IsTranslated = 1
	datastream.LineNumber = int32(parallelsentence.LineNumber)
	dstrm, err := ch.update.Update(&datastream)
	fmt.Println(dstrm) //TO-DO: this should go be front-end soon
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

// ServeHTTP implements http.Handler interface for parallel sentence create
func (ch *createHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	parallelsentenceDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	parallelsentenceDat.TranslatorID = ch.decodeContext(r)
	data, err := ch.askController(&parallelsentenceDat)

	if err != nil {
		message := "Unable to create parallelsentence error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     parallelsentence.Creater
	Update     datastream.Updater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to create parallelsentences
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create, params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.ParallelsentenceCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
