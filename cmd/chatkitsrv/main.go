package main

import (
	"github.com/gofiber/fiber"

	"github.com/nathandaly/"
)

var App = fiber.New()

func main() {
	_ = App.Listen(3000)
}

func routeHandler() {

}
