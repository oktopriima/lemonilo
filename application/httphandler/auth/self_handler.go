/*
* project lemonilo
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00:24
**/

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/application/request"
	"github.com/oktopriima/lemonilo/domain/middleware"
	"github.com/oktopriima/lemonilo/domain/response"
	"net/http"
)

func (a *authenticationHandler) SelfHandler(ctx *gin.Context) {
	ID, err := middleware.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	var req request.UserRequest
	req.ID = int(ID)

	resp, err := a.user.FindController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}

// docker run --name mysql-local --network localhost -e MYSQL_ROOT_PASSWORD=root -d mysql:latest
