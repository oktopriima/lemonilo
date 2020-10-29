/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.49
**/

package users

import (
	"github.com/jinzhu/copier"
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/core/model"
	"golang.org/x/crypto/bcrypt"
)

func (u *userController) CreateController(req request.UserRequest) (interface{}, error) {
	user := new(model.User)

	if err := copier.Copy(&user, &req); err != nil {
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)

	tx := u.db.Begin()
	defer tx.Rollback()

	m, err := u.userRepo.Create(user, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
