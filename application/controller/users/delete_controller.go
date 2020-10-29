/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:22
**/

package users

import "github.com/oktopriima/lemonilo/application/request"

func (u *userController) DeleteController(req request.UserRequest) (interface{}, error) {
	data, err := u.userRepo.Find(req.ID)
	if err != nil {
		return nil, err
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	if err = u.userRepo.Delete(data, tx); err != nil {
		return nil, err
	}

	tx.Commit()
	return data, nil
}
