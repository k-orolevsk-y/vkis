package messages

import (
	"fmt"
	"net/url"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func Features(a api.Api, object *longpoll.MessageNewObject) {
	_, _ = a.MessagesSend(url.Values{
		"peer_id": {fmt.Sprint(object.Message.PeerId)},
		"message": {"Я простой бот, который может отправить фото котиков, случайную циатату и еще немножко разного."},
	})

	return
}
