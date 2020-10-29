/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:00
**/

package auth

import (
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/config"
	"github.com/oktopriima/lemonilo/domain/core/service"
)

type AuthenticationController interface {
	LoginController(req request.LoginRequest) (interface{}, error)
}

type authenticationController struct {
	cfg         config.Config
	userService service.UserService
}

func NewAuthenticationController(cfg config.Config, userService service.UserService) AuthenticationController {
	return &authenticationController{
		cfg:         cfg,
		userService: userService,
	}
}
