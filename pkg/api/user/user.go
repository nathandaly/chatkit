package user

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"time"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	CustomData map[string]interface{} `json:"custom_data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func Routes(group *fiber.Group) {
	user := group.Group("/user")

	user.Post("/", CreateAUser)
}

func CreateAUser(c *fiber.Ctx) {
	var user User
	var result map[string]interface{}

	if err := json.Unmarshal([]byte(c.Body()), &result); err != nil {
		c.Send("Error parsing body.")
	}

	if err := c.BodyParser(&user); err != nil {
		c.Send("Error parsing user.")
	}

	user.CustomData = result["custom_data"].(map[string]interface{})
	user.DeletedAt = nil

	if err := c.JSON(user); err != nil {
		c.Send("Error converting user to JSON.")
	}
}