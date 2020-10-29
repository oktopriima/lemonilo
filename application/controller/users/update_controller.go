/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:19
**/

package users

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/oktopriima/lemonilo/application/request"
	"golang.org/x/crypto/bcrypt"
)

func (u *userController) UpdateController(req request.UserRequest) (interface{}, error) {
	data, err := u.userRepo.Find(req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed find user. err : %v", err)
	}

	if err = copier.Copy(&data, &req); err != nil {
		return nil, fmt.Errorf("failed assign request data to model. err : %v", err)
	}

	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		data.Password = string(password)
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	if err = u.userRepo.Update(data, tx); err != nil {
		return nil, err
	}

	tx.Commit()
	return data, nil
}
