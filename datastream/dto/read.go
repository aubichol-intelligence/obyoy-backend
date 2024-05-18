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
	//	ContestID string
	SourceSentence string
	LineNumber     int32
	DatasetID      string
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(delivery *model.Datastream) {
	r.SourceSentence = delivery.SourceSentence
	r.LineNumber = delivery.LineNumber
	r.DatasetID = delivery.DatasetID
}
