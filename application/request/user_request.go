/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.46
**/

package request

type UserRequest struct {
	ID       int    `uri:"id"`    // get value from uri param
	Name     string `json:"name"` // get value from form json/application
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Page     int    `form:"page"` // get value from query parameter
	Size     int    `form:"size"`
}
