package messages

import (
	"fmt"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func Previous(a api.Api, object *longpoll.MessageNewObject) {
	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {"🪄 Вернул тебя назад."},
		"keyboard": {keyboards.Start()},
	})

	return
}
