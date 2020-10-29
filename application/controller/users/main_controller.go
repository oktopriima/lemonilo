/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.45
**/

package users

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/core/service"
)

type UserController interface {
	CreateController(req request.UserRequest) (interface{}, error)
	FindController(req request.UserRequest) (interface{}, error)
	FindPagedController() (interface{}, error)
	UpdateController(req request.UserRequest) (interface{}, error)
	DeleteController(req request.UserRequest) (interface{}, error)
}

type userController struct {
	db       *gorm.DB
	userRepo service.UserService
}

func NewUserController(db *gorm.DB,
	userRepo service.UserService) UserController {
	return &userController{db, userRepo}
}
