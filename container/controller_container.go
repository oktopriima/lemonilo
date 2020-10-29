/*
* project lemonilo
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.02
**/

package container

import (
	"github.com/oktopriima/lemonilo/application/controller/auth"
	"go.uber.org/dig"

	"github.com/oktopriima/lemonilo/application/controller/users"
)

func BuildControllerContainer(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewUserController); err != nil {
		panic(err)
	}

	if err = container.Provide(auth.NewAuthenticationController); err != nil {
		panic(err)
	}

	return container
}
