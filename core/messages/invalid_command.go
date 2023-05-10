package messages

import (
	"fmt"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func InvalidCommand(a api.Api, object *longpoll.MessageNewObject) {
	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {"Не знаю такую команду, попробуй воспользоваться клавиатурой ⤵️"},
		"keyboard": {keyboards.Start()},
	})

	return
}
