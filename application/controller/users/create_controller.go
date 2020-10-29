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
	"time"

	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/core/model"
)

func (u *userController) CreateController(request request.UserRequest) (interface{}, error) {
	user := new(model.User)

	if err := copier.Copy(&user, &request); err != nil {
		return nil, err
	}
	user.LastLogin = time.Now()

	tx := u.db.Begin()
	defer tx.Rollback()

	m, err := u.userRepo.Create(user, tx)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return m, nil
}
