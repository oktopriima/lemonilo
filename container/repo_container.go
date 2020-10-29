/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.19
**/

package container

import (
	"go.uber.org/dig"

	"github.com/oktopriima/lemonilo/domain/core/service"
)

func BuildServiceContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(service.NewUserService); err != nil {
		panic(err)
	}

	return container
}
