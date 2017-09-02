package forms

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// UserSignup :
type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
Validate :
*/
func (d UserSignup) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required, validation.Length(5, 30)),
		validation.Field(&d.Mobile, validation.Required, validation.Length(13, 13)),
		validation.Field(&d.Password, validation.Required, validation.Length(5, 30)),
	)
}
