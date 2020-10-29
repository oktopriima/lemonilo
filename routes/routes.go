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
	"github.com/oktopriima/lemonilo/domain/middleware"
	"github.com/oktopriima/lemonilo/application/httphandler/users"
	"github.com/oktopriima/lemonilo/domain/config"
)

func InvokeRoute(
	engine *gin.Engine,
	user users.UserHandler,
) {
	conf := config.NewConfig()
	route := engine.Group("api/" + conf.GetString("app.version.tag") + conf.GetString("app.version.value"))

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.OPTIONS("/*path", middleware.CORSMiddleware())

	// user route
	{
		userRoute := route.Group("users")
		userRoute.POST("", user.CreateHandler)
	}

}
