package controllers

import (
	"home-provider/app"
	"home-provider/boom"
	"home-provider/forms"

	routing "github.com/go-ozzo/ozzo-routing"
)

/*
UserController :
*/
type UserController struct {
	ctx app.Context
}

/*
NewUserController :
*/
func NewUserController(ctx app.Context) *UserController {
	return &UserController{
		ctx: ctx,
	}
}

/*
CreateUser :
*/
func (u UserController) CreateUser(c *routing.Context) error {

	var form forms.UserSignup
	if err := c.Read(&form); err != nil {
		return err
	}

	if err := form.Validate(); err != nil {
		return err
	}
	//return c.Write(json)

	return boom.BadRequest("I am bad request")
}
