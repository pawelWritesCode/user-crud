package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	container := newContainer()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/alive", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]interface{}{"status": "alive"})
	})

	//Users CRUD
	e.GET("/users/:userId", container.ControllerUser.GetSingle)
	e.GET("/users", container.ControllerUser.GetMany)
	e.POST("/users", container.ControllerUser.Create)
	e.PUT("/users/:userId", container.ControllerUser.Replace)
	e.DELETE("/users/:userId", container.ControllerUser.Delete)
	e.POST("/users/:userId/avatar", container.ControllerUser.ReceiveAvatar)

	fmt.Println("Available routes:")
	routes := e.Routes()
	for _, route := range routes {
		fmt.Printf("%s \t %s\n", route.Method, route.Path)
	}

	e.Logger.Fatal(e.Start(":1234"))
}
