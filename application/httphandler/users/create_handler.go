/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.05
**/

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/response"
)

func (u *userHandler) CreateHandler(ctx *gin.Context) {
	var req request.UserRequest

	if err := ctx.ShouldBindJSON(&ctx); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := u.user.CreateController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response.NewSuccessResponse(ctx, result)
}
