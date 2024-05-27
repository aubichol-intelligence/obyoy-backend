package dto

import "obyoy-backend/model"

// ReadResp holds the response data for reading user
type ReadResp struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	//
	ID string `json:"id"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(user *model.User) {
	r.Name = user.FirstName
	r.Email = user.Email
	r.Address = user.Address
	r.PhoneNumber = user.PhoneNumber
	//
	r.ID = user.ID
}
