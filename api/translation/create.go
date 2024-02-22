package translation

import (
	"io"
	"net/http"

	"horkora-backend/api/middleware"
	"horkora-backend/api/routeutils"
	"horkora-backend/apipattern"
	"horkora-backend/translation"
	"horkora-backend/translation/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// createHandler holds handler for creating translation items
type createHandler struct {
	create translation.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	translation dto.translation,
	err error,
) {
	err = translation.FromReader(body)
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
	translation *dto.translation,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(translation)
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

	translationDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	translationDat.UserID = ch.decodeContext(r)
	data, err := ch.askController(&translationDat)

	if err != nil {
		message := "Unable to create translation error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     translation.Creater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to take translations
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.translationCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}