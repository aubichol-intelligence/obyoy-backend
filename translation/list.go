package translation

import (
	"obyoy-backend/errors"
	"obyoy-backend/model"
	storetranslation "obyoy-backend/store/translation"
	"obyoy-backend/translation/dto"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Reader provides an interface for reading translationes
type Lister interface {
	List(req *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error)
}

// translationReader implements Reader interface
type translationLister struct {
	translations storetranslation.Translations
}

func (list *translationLister) askStore(state string, skip int64, limit int64) (
	translation []*model.Translation,
	err error,
) {
	translation, err = list.translations.FindByTranslationID(state, skip, limit)
	return
}

func (list *translationLister) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (list *translationLister) prepareResponse(
	translations []*model.Translation,
) (
	resp []dto.ReadResp,
) {
	for _, translation := range translations {
		var tmp dto.ReadResp
		tmp.FromModel(translation)
		resp = append(resp, tmp)
	}
	//resp.FromModel(translation)
	return
}

func (read *translationLister) List(translationReq *dto.ListReq, skip int64, limit int64) ([]dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	translations, err := read.askStore(translationReq.UserID, skip, limit)
	if err != nil {
		logrus.Error("Could not find translation error : ", err)
		return nil, read.giveError()
	}

	var resp []dto.ReadResp
	resp = read.prepareResponse(translations)

	return resp, nil
}

// NewReaderParams lists params for the NewReader
type NewListerParams struct {
	dig.In
	Translation storetranslation.Translations
}

// NewReader provides Reader
func NewList(params NewReaderParams) Lister {
	return &translationLister{
		translations: params.Translation,
	}
}
