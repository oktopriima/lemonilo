/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 21:36
 * Copyright (c) 2020
 */

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/lemonilo/domain/config"
)

type ServerRoute struct {
	cfg    config.Config
	engine *gin.Engine
}

func NewRoute(cfg config.Config, engine *gin.Engine) *ServerRoute {
	return &ServerRoute{cfg, engine}
}

func (s *ServerRoute) Run() {
	if err := s.engine.Run(fmt.Sprintf(":%s", s.cfg.GetString("server.address"))); err != nil {
		log.Fatal(err)
	}
}
