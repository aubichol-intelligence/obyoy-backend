package translation

import (
	"io"
	"net/http"

	"obyoy-backend/api/middleware"
	"obyoy-backend/api/routeutils"
	"obyoy-backend/apipattern"
	"obyoy-backend/translation"
	"obyoy-backend/translation/dto"

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
	translation dto.Create,
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
	translation *dto.Create,
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

// ServeHTTP implements http.Handler interface for creating translation
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

	//	translationDat.UserID = ch.decodeContext(r)
	data, err := ch.askController(&translationDat)

	if err != nil {
		message := "Unable to create translation error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

// CreateParams provide parameters for CreateRoute for translation
type CreateParams struct {
	dig.In
	Create     translation.Creater
	Middleware *middleware.Auth
}

// CreateRoute provides a route that lets to create translations
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.TranslationCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
