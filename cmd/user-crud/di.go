package main

import (
	"github.com/pawelWritesCode/user-crud/internal/controller"
	"github.com/pawelWritesCode/user-crud/internal/database"
	"github.com/pawelWritesCode/user-crud/internal/manager"
)

type Container struct {
	ControllerUser controller.User
}

func newContainer() *Container {
	db := database.New()
	um := manager.NewUsersManager(db)
	controllerUser := controller.NewUser(um)

	return &Container{ControllerUser: controllerUser}
}
