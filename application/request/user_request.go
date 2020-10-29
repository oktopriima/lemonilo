/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.46
**/

package request

type UserRequest struct {
	ID       int    `uri:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
