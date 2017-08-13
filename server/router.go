package server

import (
	"home-provider/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := router.Group("v1")
	{
		// init user controller
		user := new(controllers.UserController)

		// create user routes
		usersRoutes := v1.Group("users")
		usersRoutes.GET("/", user.CreateUser)
	}
}
