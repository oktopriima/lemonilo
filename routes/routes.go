/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 21:33
 * Copyright (c) 2020
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/application/httphandler/auth"
	"github.com/oktopriima/lemonilo/application/httphandler/users"
	"github.com/oktopriima/lemonilo/domain/middleware"
)

func InvokeRoute(
	engine *gin.Engine,
	user users.UserHandler,
	auth auth.AuthenticationHandler,
) {
	route := engine.Group("lemonilo-api/")

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.OPTIONS("/*path", middleware.CORSMiddleware())

	// user route
	{
		userRoute := route.Group("users")
		userRoute.POST("", user.CreateHandler)
		userRoute.GET(":id", user.FindHandler)
		userRoute.GET("", user.FindPagedHandler)
		userRoute.PUT(":id", user.UpdateHandler)
		userRoute.DELETE(":id", user.DeleteHandler)
	}

	// login
	{
		loginRoute := route.Group("auth")
		loginRoute.POST("", auth.LoginHandler)
	}

	// endpoint with token
	{
		authRoute := route.Group("me")
		authRoute.Use(middleware.MyAuth("app"))
		authRoute.GET("", user.FindHandler)
	}

}
