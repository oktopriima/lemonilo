/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:17
**/

package users

import "github.com/oktopriima/lemonilo/application/request"

func (u *userController) FindPagedController(req request.UserRequest) (interface{}, error) {
	// example of criteria
	// criteria := make(map[string]interface{})
	// criteria["email"] = req.Email

	data, err := u.userRepo.FindBy(nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}
