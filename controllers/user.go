package controllers

import (
	"home-provider/app"
	"home-provider/forms"
	"home-provider/models"

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
func (ctrl UserController) CreateUser(c *routing.Context) error {

	var form forms.UserSignup
	if err := c.Read(&form); err != nil {
		return err
	}

	if err := form.Validate(); err != nil {
		return err
	}

	user := models.User{}
	userDao := ctrl.ctx.Db().GetUserDao()
	if err := userDao.CreateUser(&user); err != nil {
		return err
	}

	return c.Write(user)
}
