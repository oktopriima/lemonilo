/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:52
 * Copyright (c) 2020
 */

package helper

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/lemonilo/domain/config"
	"github.com/oktopriima/lemonilo/domain/response"
	"net/http"
)

func Localization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := config.NewConfig()
		session := sessions.Default(ctx)

		var lang string
		v := session.Get("lang")
		if ctx.Request.Header.Get("Accept-Language") != "" {
			lang = ctx.Request.Header.Get("Accept-Language")
		} else {
			if v == nil {
				lang = conf.GetString("app.lang")
			} else {
				lang = v.(string)
			}
		}

		session.Set("lang", lang)
		err := session.Save()
		if err != nil {
			abortMission(ctx, err)
		}

		if err := InitLocales("lang", lang); err != nil {
			abortMission(ctx, err)
		}

		ctx.Next()
	}
}

func abortMission(ctx *gin.Context, err error) {
	resp := new(response.Error)
	resp.ErrorCode = http.StatusUnauthorized
	resp.Message = err.Error()
	resp.Status = http.StatusUnauthorized

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, &resp)
	return
}
