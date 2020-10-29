/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:03
**/

package auth

import (
	"fmt"
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (a *authenticationController) LoginController(req request.LoginRequest) (interface{}, error) {
	user, err := a.userService.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed find user. err : %v", err)
	}

	byteHash := []byte(user.Password)
	bytePlain := []byte(req.Password)
	if err = bcrypt.CompareHashAndPassword(byteHash, bytePlain); err != nil {
		return nil, fmt.Errorf("password didnt match. err : %v", err)
	}

	m := middleware.NewCustomAuth([]byte(a.cfg.GetString("app.signature")))
	token, err := m.GenerateToken(middleware.TokenStructure{
		UserID: int64(user.ID),
		Email:  user.Email,
		Role:   "app",
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
