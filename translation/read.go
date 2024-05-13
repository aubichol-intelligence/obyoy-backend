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
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

// translationReader implements Reader interface
type translationReader struct {
	translations storetranslation.Translations
}

func (read *translationReader) askStore(translationID string) (
	translation *model.Translation,
	err error,
) {
	translation, err = read.translations.FindByID(translationID)
	return
}

func (read *translationReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *translationReader) prepareResponse(
	translation *model.Translation,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(translation)
	return
}

func (read *translationReader) Read(translationReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	translation, err := read.askStore(translationReq.TranslationID)
	if err != nil {
		logrus.Error("Could not find translation error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(translation)

	return &resp, nil
}

// NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Translation storetranslation.Translations
}

// NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &translationReader{
		translations: params.Translation,
	}
}
