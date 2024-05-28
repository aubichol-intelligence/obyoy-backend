package dto

import "obyoy-backend/model"

// ReadReq stores translation read request data
type ReadReq struct {
	UserID        string
	TranslationID string
}

// ReadReq stores translation read request data
type ReadResp struct {
	ID                  string `json:"translation_id"`
	DatasetID           string `json:"dataset_id"`
	DatastreamID        string `json:"datastream_id"`
	TranslatorID        string `json:"translator_id"`
	ReviewerID          string `json:"reviewer_id"`
	SourceSentence      string `json:"source_sentence"`
	SourceLanguage      string `json:"source_language"`
	DestinationSentence string `json:"destination_sentence"`
	DestinationLanguage string `json:"destination_language"`
	LineNumber          int    `json:"line_number"`
	Name                string `json:"name"`
}

// FromModel converts the translation model data to response data
func (r *ReadResp) FromModel(translation *model.Translation) {
	r.ID = translation.ID
	r.DatasetID = translation.DatasetID
	r.DatastreamID = translation.DatastreamID
	r.TranslatorID = translation.TranslatorID
	r.ReviewerID = translation.ReviewerID
	r.SourceLanguage = translation.SourceLanguage
	r.SourceSentence = translation.SourceSentence
	r.DestinationSentence = translation.DestinationSentence
	r.DestinationLanguage = translation.DestinationLanguage
	r.LineNumber = translation.LineNumber
	r.Name = translation.Name
}
