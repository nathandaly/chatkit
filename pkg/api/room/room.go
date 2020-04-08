package room

import (
	"github.com/gofiber/fiber"
	"time"
)

type Membership struct {
	RoomId string `json:"room_id"`
	UserIds map[string]string `json:"user_ids"`
}

type Room struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatedById string `json:"created_by_id"`
	CustomData map[string]interface{} `json:"custom_data"`
	Private bool `json:"private"`
	LastMessageAt time.Time `json:"last_message_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Membership Membership `json:"membership"`
}

func Routes(group *fiber.Group) {
	room := group.Group("/room")

	room.Get("/", GetAllRooms)
	room.Get("/:roomID", GetARoom)
}

// Query Parameters:
//		from_id (string|optional): ID (exclusive) from which rooms with larger IDs
//		should be returned. The number of room returned is at most 100 per request.
//		include_private (string|optional): If `true` will also return private rooms
//		present in the instance. This requires a su token.
// Response codes:
//		200	400	404	422	503
func GetAllRooms(ctx *fiber.Ctx)  {
	ctx.Send("GetAllRooms")
}

// Response codes:
//		200	422	503
func GetARoom(ctx *fiber.Ctx)  {
	ctx.Send("GetARoom")
}