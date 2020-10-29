/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:48
 * Copyright (c) 2020
 */

package container

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"

	"github.com/oktopriima/lemonilo/domain/config"
)

func BuildConfigContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() config.Config {
		return config.NewConfig()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func(cfg config.Config) (error, *gorm.DB) {
		return config.NewMysqlConfig(cfg)
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		panic(err)
	}

	return container
}
