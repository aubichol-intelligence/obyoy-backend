package dto

import "obyoy-backend/model"

// ReadReq stores order read request data
type ReadReq struct {
	UserID       string
	DatastreamID string
}

// ReadReq stores order read request data
type ReadResp struct {
	DatastreamID   string `json:"datastream_id"`
	SourceSentence string `json:"source_sentence"`
	SourceLanguage string `json:"source_language"`
	LineNumber     int32  `json:"line_number"`
	DatasetID      string `json:"dataset_id"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(delivery *model.Datastream) {
	r.SourceSentence = delivery.SourceSentence
	r.SourceLanguage = delivery.SourceLanguage
	r.LineNumber = delivery.LineNumber
	r.DatasetID = delivery.DatasetID
	r.DatastreamID = delivery.ID
}
