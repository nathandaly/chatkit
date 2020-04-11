## Chatkit Message Package

#### Todo

###### [Message entity](https://pusher.com/docs/chatkit/reference/latest#message-entity)

- `id`: The message ID. Message IDs are integers, and increase over time. They determine the canonical ordering of messages within a room.
- `user_id`: The ID of the user who sent the message.
- `room_id`: The ID of the room to which the message belongs.
- `created_at`: The creation time of the message.updated_at: The time of the last update to the message.
- `parts`: an array of the message’s parts.


- [ ] [Persistence](https://medium.com/@shanev/simple-and-easy-data-persistence-in-go-aa019a6f3106)
- [x]  [Create a user](https://pusher.com/docs/chatkit/reference/latest#create-a-user)
- [x]  [Batch create users](https://pusher.com/docs/chatkit/reference/latest#batch-create-users)
- [x]  [Get user](https://pusher.com/docs/chatkit/reference/latest#get-user)
- [x]  [Get users](https://pusher.com/docs/chatkit/reference/latest#get-users)
- [ ]  [Get users by IDs](https://pusher.com/docs/chatkit/reference/latest#get-users-by-ids)
- [ ]  [Update a user](https://pusher.com/docs/chatkit/reference/latest#update-a-user)
- [ ]  [Delete a user](https://pusher.com/docs/chatkit/reference/latest#delete-a-user)

###### [Part entity](https://pusher.com/docs/chatkit/reference/latest#part-entity)

> Parts hold the content for a message. There are three variants of parts:

 - `inline` parts store and deliver their content directly as part of the message entity whenever it is stored or retrieved. When the size of the piece of content does not exceed their limit, this is the most convenient way to deliver content.
 - `link` parts store a URL which references some content hosted publically elsewhere, for example in a cloud storage bucket associated with your application.
 - `attachment` parts represent a piece of content which has been uploaded and stored by Chatkit separately from the message itself (probably because it exceeds the size limit for inline content).

> All part entities have one field in common:

 - `type` The [MIME type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types) of the content which the part holds.

> Each part variant contains exactly one of the following fields:

 - `content`, for `inline` parts. This contains the actual content of the part.
 - `url`, for `link` parts. This is the URL at which the actual content can be accessed.
 - `attachment`, for `attachment` parts. This entity describes how to access a piece of content which is stored by Chatkit. The entity fields are described directly below.
 
###### [Attachment entity](https://pusher.com/docs/chatkit/reference/latest#attachment-entity)
 
 > An Attachment describes a piece of content stored by Chatkit. It has the following fields.

 - `id` The unique ID of the attachment.
 - `download_url` Signed URL at which the content may be retrieved. The signature on this URL expires after a certain period of time.
 - `refresh_url` When the signature on the download URL has expired, the entity may be refreshed at this URL in order to get an up to date signature.
 - `expiration` Timestamp at which time the download_url will expire.
 - `name` The filename associated with the content.
 - `size` The size of the content in bytes.
 - `custom_data` Any custom data that was associated with the attachment when it was created.
 
###### [Example complete message entity](https://pusher.com/docs/chatkit/reference/latest#example-complete-message-entity)

```json
{
  "id": 45617234,
  "user_id": "bob",
  "room_id": "538a8fc",
  "parts": [
    {
      "type": "text/plain",
      "content": "see attached document and figure"
    },
    {
      "type": "image/png",
      "url": "https://example.com/figure01.png"
    },
    {
      "type": "application/pdf",
      "attachment": {
        "id": "793c8a94-1702-4b7b-92aa-62f3be1a8efb",
        "download_url": "https://example.download.url/793c8a94-1702-4b7b-92aa-62f3be1a8efb",
        "refresh_url": "https://example.refresh.url/793c8a94-1702-4b7b-92aa-62f3be1a8efb",
        "expiration": "2019-02-12T17:29:22Z",
        "name": "super-interesting-document.pdf",
        "size": 4096,
        "custom_data": {
          "foo": "bar"
        }
      }
    }
  ],
  "created_at": "2017-03-23T11:36:42Z",
  "updated_at": "2017-03-23T11:36:42Z"
}
```

 - [ ] [Send a message](https://pusher.com/docs/chatkit/reference/latest#send-a-message)
 - [ ] [Upload an attachment](https://pusher.com/docs/chatkit/reference/latest#upload-an-attachment)
 - [ ] [Edit a message](https://pusher.com/docs/chatkit/reference/latest#edit-a-message)
 - [ ] [Get messages from a room](https://pusher.com/docs/chatkit/reference/latest#get-messages-from-a-room)
 - [ ] [Get a single message from a room](https://pusher.com/docs/chatkit/reference/latest#get-a-single-message-from-a-room)
 - [ ] [Delete a message](https://pusher.com/docs/chatkit/reference/latest#delete-a-message)