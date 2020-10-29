/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:22
**/

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/response"
	"net/http"
)

func (a *authenticationHandler) LoginHandler(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.auth.LoginController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}
