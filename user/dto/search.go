package dto

import "obyoy-backend/model"

// Search stores search related data
type Search struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
}

// FromModel converts model data to json type data
func (s *Search) FromModel(user *model.User) {
	s.ID = user.ID
	s.FirstName = user.FirstName
	s.LastName = user.LastName
	s.Gender = user.Gender
	s.Email = user.Email
}
