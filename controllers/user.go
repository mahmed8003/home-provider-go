package controllers

import (
	"home-provider/app"
	"net/http"
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
func (u UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	/*
		var json forms.UserSignup
		if c.BindJSON(&json) == nil {
			c.JSON(http.StatusOK, json)
		} else {
			//c.JSON(406, gin.H{"message": "Invalid form", "form": json})
			//c.Abort()
		}
	*/
	w.Write([]byte("hi"))

}
