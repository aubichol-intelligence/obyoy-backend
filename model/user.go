package model

import "time"

// User defines user model
type User struct {
	ID           string
	FirstName    string
	LastName     string
	Gender       string
	Address      string
	BirthDate    BirthDate
	Email        string
	Password     string
	Verified     bool
	Profile      UserDetails
	ProfilePic   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Suspended    bool
	IsDriver     bool
	UserType     string
	PhoneNumber  string
	AccountType  string
	Latitude     float64
	Longitude    float64
	RestaurantID string
}

// BirthDate defines birthdate model
type BirthDate struct {
	Year  int
	Month int
	Day   int
}

// TO-DO: need to decide the details later
type UserDetails struct {
	Year  int
	Month int
	Day   int
}
