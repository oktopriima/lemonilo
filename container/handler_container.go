/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.11
**/

package container

import (
	"go.uber.org/dig"

	"github.com/oktopriima/lemonilo/application/httphandler/users"
)

func BuildHandlerContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserHandler); err != nil {
		panic(err)
	}

	return container
}
