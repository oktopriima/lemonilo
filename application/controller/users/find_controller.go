/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:13
**/

package users

import (
	"fmt"
	"github.com/oktopriima/lemonilo/application/request"
)

func (u *userController) FindController(req request.UserRequest) (interface{}, error) {
	data, err := u.userRepo.Find(req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed find user. err : %v", err)
	}

	return data, nil
}
