package messages

import (
	"fmt"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func Functions(a api.Api, object *longpoll.MessageNewObject) {
	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {"🍃 Я отправил тебе список функций в клавитауре."},
		"keyboard": {keyboards.Functions()},
	})

	return
}
