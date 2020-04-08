package user

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func Routes(group *fiber.Group) {
	user := group.Group("/user")

	user.Get("/:name", func(c *fiber.Ctx) {
		c.Send("Hello, World!" + c.Params("name"))
		fmt.Printf("Hello %s!", c.Params("name"))
		// => Hello john!
	})
}