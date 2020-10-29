/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 21:12
 * Copyright (c) 2020
 */

package main

import (
	"go.uber.org/dig"

	"github.com/oktopriima/lemonilo/container"
)

func BuildContainer() *dig.Container {
	c := dig.New()
	c = container.BuildConfigContainer(c)
	c = container.BuildControllerContainer(c)
	c = container.BuildHandlerContainer(c)
	c = container.BuildServiceContainer(c)

	return c
}
