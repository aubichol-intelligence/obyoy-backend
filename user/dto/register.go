package dto

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"

	"obyoy-backend/errors"

	validator "gopkg.in/go-playground/validator.v9"
)

// Register provides dto for user register
type Register struct {
	//	FirstName   string    `json:"first_name" validate:"required,min=2,max=20"`
	FirstName string `json:"first_name"`
	//	LastName    string    `json:"last_name" validate:"required,min=2,max=20"`
	LastName string `json:"last_name"`
	//	Gender      string    `json:"gender" validate:"required,eq=male|eq=female|eq=other"`
	Gender    string    `json:"gender"`
	BirthDate BirthDate `json:"birth_date"`
	//	Email       string    `json:"email" validate:"required,email"`
	Email       string `json:"email"`
	Password    string `json:"password" validate:"required,min=6"`
	Suspended   bool   `json:"suspended"`
	IsDriver    bool   `json:"is_driver"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	AccountType string `json:"account_type" validate:"required,eq=driver|eq=admin|eq=restaurant"`
}

// Validate validates registration data
func (r *Register) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			errors.Base{"invalid data", false},
		})
	}
	return nil
}

func (r *Register) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(r)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid register data", false},
		})
	}

	if r.Password == "" {
		fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"empty password", false},
		})
	}

	h := md5.New()
	_, err = io.WriteString(h, r.Password)
	if err != nil {
		fmt.Errorf("%s:%w", err.Error(), &errors.Unknown{
			Base: errors.Base{"could not convert password to hash", false},
		})
	}
	r.Password = fmt.Sprintf("%x", h.Sum(nil))

	return nil
}
