## Chatkit Room Package

#### Todo

###### [Room entity](https://pusher.com/docs/chatkit/reference/latest#room-entity)
```json
{
  "id": "ac43dfef",
  "name": "Chatkit chat",
  "created_by_id": "alice",
  "push_notification_title_override": null,
  "custom_data": {
    "highlight_color": "blue"
  },
  "private": false,
  "last_message_at": "2020-01-08T14:55:10Z",
  "created_at": "2017-03-23T11:36:42Z",
  "updated_at": "2017-03-23T11:36:42Z"
}
```
- `id` (string): A unique id for the room.
- `created_by_id` (string): User id that created the room.
- `name` (string): Room name / title.
- `push_notification_title_override` (string|optional): Title to be used for push notifications.
- `private` (boolean): Indicates if the room is private or public.
- `custom_data` (object): Custom data associated with the entity.
- `last_message_at` (string|optional): Timestamp of the last message in the room, if there is one.
- `created_at` (string): Timestamp of the entity creation.
- `updated_at` (string): The last time the entity was updated.
- `deleted_at` (string|optional): When the entity was marked as deleted.

###### [Membership entity](https://pusher.com/docs/chatkit/reference/latest#membership-entity)
```json
{
  "room_id": "ac43dfef",
  "user_ids": [
    "alice",
    "carol"
  ]
}
```
- `room_id` (string): The id of the Room these are the members for.
- `user_ids` (array) A list of the User ids which are members of the room.

###### [Read State entity](https://pusher.com/docs/chatkit/reference/latest#read-state-entity)
```json
{
  "room_id": "ac43dfef",
  "unread_count": 3,
  "cursor": {
    "room_id": "ac43dfef",
    "user_id": "alice",
    "position": 43398,
    "cursor_type": 0,
    "updated_at": "2020-01-03T18:15:15Z"
  }
}
```
- `room_id` (string): The id of the Room this read state applies to.
- `unread_count` (integer): The number of unread messages in the room.
- `cursor` (Cursor|optional): The cursor used to calculate the unread count.


- [ ] [Persistence](https://medium.com/@shanev/simple-and-easy-data-persistence-in-go-aa019a6f3106)
- [ ]  [Create a room](https://pusher.com/docs/chatkit/reference/latest#create-a-room)
- [ ]  [Get a room](https://pusher.com/docs/chatkit/reference/latest#get-a-room)
- [ ]  [Get rooms](https://pusher.com/docs/chatkit/reference/latest#get-rooms)
- [ ]  [Delete a room](https://pusher.com/docs/chatkit/reference/latest#delete-a-room)
- [ ]  [Update a room](https://pusher.com/docs/chatkit/reference/latest#update-a-room)
- [ ]  [Add users](https://pusher.com/docs/chatkit/reference/latest#add-users)
- [ ]  [Remove users](https://pusher.com/docs/chatkit/reference/latest#remove-users)
- [ ]  [Typing indicators](https://pusher.com/docs/chatkit/reference/latest#typing-indicators)