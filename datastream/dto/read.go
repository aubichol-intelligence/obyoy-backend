package dto

import "obyoy-backend/model"

// ReadReq stores datastream read request data
type ReadReq struct {
	UserID       string
	DatastreamID string
}

// ReadReq stores datastream read request data
type ReadResp struct {
	DatastreamID   string `json:"datastream_id"`
	SourceSentence string `json:"source_sentence"`
	SourceLanguage string `json:"source_language"`
	LineNumber     int32  `json:"line_number"`
	DatasetID      string `json:"dataset_id"`
	Name           string `json:"name"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(datastream *model.Datastream) {
	r.SourceSentence = datastream.SourceSentence
	r.SourceLanguage = datastream.SourceLanguage
	r.LineNumber = datastream.LineNumber
	r.DatasetID = datastream.DatasetID
	r.DatastreamID = datastream.ID
	r.Name = datastream.Name
}
