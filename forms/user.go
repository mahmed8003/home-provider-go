package forms

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// UserSignup :
type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	BirthDay string `json:"birthday" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	PhotoURL string `json:"photo_url" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

/*
Validate :
*/
func (d UserSignup) Validate() error {
	return validation.ValidateStruct(&d,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&d.Name, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&d.BirthDay, validation.Required, validation.Length(5, 50)),
	)
}
