package controllers

import (
	"home-provider/app"
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
	c.String(http.StatusOK, "pong1")
}
