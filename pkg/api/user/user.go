package user

import (
	"fmt"
	"github.com/gofiber/fiber"
	"time"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	CustomData string `json:"custom_data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func Routes(group *fiber.Group) {
	user := group.Group("/user")

	user.Post("/", CreateAUser)
}

func CreateAUser(c *fiber.Ctx) {
	c.Send("Hello, World!" + c.Params("name"))
	fmt.Printf("Hello %s!", c.Params("name"))
	// => Hello john!
}