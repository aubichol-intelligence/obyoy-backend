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
func (r *ReadResp) FromModel(delivery *model.Translation) {
	r.ID = delivery.ID
	r.DatasetID = delivery.DatasetID
	r.DatastreamID = delivery.DatastreamID
	r.TranslatorID = delivery.TranslatorID
	r.ReviewerID = delivery.ReviewerID
	r.SourceLanguage = delivery.SourceLanguage
	r.SourceSentence = delivery.SourceSentence
	r.DestinationSentence = delivery.DestinationSentence
	r.DestinationLanguage = delivery.DestinationLanguage
	r.LineNumber = delivery.LineNumber
	r.Name = delivery.Name
}
