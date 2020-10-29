/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:49
**/

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/response"
	"net/http"
)

func (u *userHandler) FindHandler(ctx *gin.Context) {
	var req request.UserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := u.user.FindController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}
