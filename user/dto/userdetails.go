package dto

import "obyoy-backend/model"

// BirthDate stores birth data data
type UserDetails struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// FromModel converts model data to json type data
func (b *UserDetails) FromModel(data *model.UserDetails) {
	b.Year = data.Year
	b.Month = data.Month
	b.Day = data.Day
}

// ToModel converts json type data to model data
func (b *UserDetails) ToModel(m *model.UserDetails) {
	m.Year = b.Year
	m.Month = b.Month
	m.Day = b.Day
}
