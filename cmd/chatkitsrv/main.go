package main

import (
	"../../pkg/user"
	"github.com/gofiber/fiber"
)

var App = fiber.New()

func main() {
	api := App.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/user", user.Routes)

	_ = App.Listen(3000)
}
