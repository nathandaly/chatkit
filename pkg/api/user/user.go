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
	CustomData map[string]interface{} `json:"custom_data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func Routes(group *fiber.Group) {
	user := group.Group("/user")

	user.Get("/", GetAllUsers)
	user.Get("/:userID", GetAUser)
	user.Post("/", CreateAUser)
	user.Post("/batch_users", CreateBatchUsers)
}

func GetAllUsers(c *fiber.Ctx) {
	var users []User

	if err := c.JSON(users); err != nil {
		c.Send("Error converting user collections to JSON.")
	}
}

func GetAUser(c *fiber.Ctx) {
	user := User{
		Id: "alice",
		Name: "Alice A",
		AvatarUrl: "https://example.com/img/alice.png",
		CustomData: map[string]interface{}{
			"email": "alice@example.com",
			"age": "21",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	if err := c.JSON(user); err != nil {
		c.Send("Error converting user to JSON.")
	}
}

func CreateAUser(c *fiber.Ctx) {
	var user User

	if err := c.BodyParser(&user); err != nil {
		c.Send("Error parsing user.")
	}

	if err := c.JSON(user); err != nil {
		c.Send("Error converting user to JSON.")
	}
}

func CreateBatchUsers(c *fiber.Ctx) {
	var users []User

	if err := c.BodyParser(&users); err != nil {
		c.Send("Error parsing users.")
	}

	for _, user := range users {
		fmt.Println("User:", user.CustomData)
		_ = createUser(user)
	}

	if err := c.JSON(users); err != nil {
		c.Send("Error converting users to JSON.")
	}
}

func createUser(user User) User {
	user.DeletedAt = nil

	// Persist use to storage.

	return user
}