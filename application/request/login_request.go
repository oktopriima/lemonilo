/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:01
**/

package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
