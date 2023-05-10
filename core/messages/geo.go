package messages

import (
	"fmt"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func Geo(a api.Api, object *longpoll.MessageNewObject) {
	if !object.Message.Geo.Has() {
		_, _ = a.MessagesSend(url.Values{
			"peer_id":  {fmt.Sprint(object.Message.PeerId)},
			"message":  {"🔴 Мне не удалось получить информацию о твоей геолокации :("},
			"keyboard": {keyboards.RemoveMessage()},
		})
		return
	}

	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {fmt.Sprintf("🌏 Ого! Ты проживаешь в %v. Я удивлён.", object.Message.Geo.Place.Title)},
		"keyboard": {keyboards.RemoveMessage()},
	})
}
