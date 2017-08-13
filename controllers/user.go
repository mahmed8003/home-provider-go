package controllers

import (
	"github.com/kataras/iris/context"
)

/*
UserController :
*/
type UserController struct{}

/*
CreateUser :
*/
func (u UserController) CreateUser(ctx context.Context) {
	ctx.WriteString("pong")
	println("Hello dellop")
}
