package controllers

import (
	"home-provider/app"
	"home-provider/forms"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (u UserController) CreateUser(c *gin.Context) {
	var json forms.UserSignup
	if c.BindJSON(&json) == nil {
		c.JSON(http.StatusOK, json)
	} else {
		//c.JSON(406, gin.H{"message": "Invalid form", "form": json})
		//c.Abort()
	}

}
