/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 21:14
 * Copyright (c) 2020
 */

package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oktopriima/lemonilo/domain/config"
	"github.com/oktopriima/lemonilo/domain/middleware"
	"github.com/oktopriima/lemonilo/routes"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	var err error

	c := BuildContainer()

	cfg := config.NewConfig()
	if err, _ = middleware.NewMiddlewareConfig(cfg); err != nil {
		panic(err)
	}

	if err = c.Invoke(routes.InvokeRoute); err != nil {
		panic(err)
	}

	if err = c.Provide(NewRoute); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(s *ServerRoute) {
		s.Run()
	}); err != nil {
		panic(err)
	}
}
