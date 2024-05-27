package dto

import "obyoy-backend/model"

// ReadReq stores order read request data
type ReadReq struct {
	UserID       string
	DatastreamID string
}

// ReadReq stores order read request data
type ReadResp struct {
	//	UserID    string
	//	datastreamID string
	SourceSentence string `json:"source_sentence"`
	LineNumber     int32  `json:"line_number"`
	DatasetID      string `json:"dataset_id"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(delivery *model.Datastream) {
	r.SourceSentence = delivery.SourceSentence
	r.LineNumber = delivery.LineNumber
	r.DatasetID = delivery.DatasetID
}
