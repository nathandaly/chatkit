package main

import (
	"../../pkg/api/room"
	"../../pkg/api/user"
	"../../routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
)

func main() {
	app := fiber.New(&fiber.Settings{
		TemplateFolder:    "../../views",
		TemplateExtension: ".html",
	})

	app.Use(recover.New(recover.Config{
			Handler: func(c *fiber.Ctx, err error) {
			c.Status(500)
			c.Render("500", fiber.Map{
					"Title":   "500 Bad Request - Fiber",
					"Message": err.Error(),
			})
		},
	}))

	app.Use(logger.New())
	app.Use(helmet.New())

	api := app.Group("/api")
	user.Routes(api)
	room.Routes(api)

	app.Static("/", "./public")

	app.Use(routes.NotFound)

	app.Listen(3000)
}
