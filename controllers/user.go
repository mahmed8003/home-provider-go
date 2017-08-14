package controllers

import (
	"home-provider/app"
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

	/*
		var json forms.UserSignup
		if c.BindJSON(&json) == nil {
			c.JSON(http.StatusOK, json)
		} else {
			//c.JSON(406, gin.H{"message": "Invalid form", "form": json})
			//c.Abort()
		}
	*/
	var json forms.UserSignup
	return c.Write(json)
}
