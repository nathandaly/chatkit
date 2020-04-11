package message

import "github.com/gofiber/fiber"

type message struct {

}

func Routes(group *fiber.Group) {
	message := group.Group("/message")

	message.Post("/rooms/:roomID/messages", SendAMessage)
}

func SendAMessage(ctx *fiber.Ctx)  {

}