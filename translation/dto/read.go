package dto

import "obyoy-backend/model"

// ReadReq stores order read request data
type ReadReq struct {
	UserID        string
	TranslationID string
}

// ReadReq stores order read request data
type ReadResp struct {
	UserID    string
	ContestID string
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(delivery *model.Translation) {
	// r.Name = delivery.Name
	// r.Phone = delivery.Phone
	// r.Address = delivery.Address
	// r.UserID = delivery.UserID
}
