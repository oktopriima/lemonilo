/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:20
**/

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/application/controller/auth"
	"github.com/oktopriima/lemonilo/application/controller/users"
)

type AuthenticationHandler interface {
	LoginHandler(ctx *gin.Context)
	SelfHandler(ctx *gin.Context)
}

type authenticationHandler struct {
	auth auth.AuthenticationController
	user users.UserController
}

func NewAuthenticationHandler(auth auth.AuthenticationController,
	user users.UserController) AuthenticationHandler {
	return &authenticationHandler{auth: auth, user: user}
}
