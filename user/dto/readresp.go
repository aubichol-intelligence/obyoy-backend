package dto

import "obyoy-backend/model"

// ReadResp holds the response data for reading restaurant
type ReadResp struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	//
	ID string `json:"id"`
}

// FromModel converts the model data to response data
func (r *ReadResp) FromModel(restaurant *model.User) {
	r.Name = restaurant.FirstName
	r.Email = restaurant.Email
	r.Address = restaurant.Address
	r.PhoneNumber = restaurant.PhoneNumber
	//
	r.ID = restaurant.ID
}
