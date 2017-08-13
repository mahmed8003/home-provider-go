package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
UserController :
*/
type UserController struct{}

/*
CreateUser :
*/
func (u UserController) CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "pong1")
	println("Hello dellop")
}
