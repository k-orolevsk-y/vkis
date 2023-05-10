package core

import (
	"fmt"
	"log"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/core/messages"
	"vkIntership/vk/longpoll"
)

func (c Core) messagesNew(object *longpoll.MessageNewObject) {
	a := c.bot.GetApi()

	if object.Message.Payload != "" {
		var payload keyboards.Payload
		if err := object.Message.GetPayload(&payload); err != nil {
			_, _ = a.MessagesSend(url.Values{
				"peer_id": {fmt.Sprint(object.Message.PeerId)},
				"message": {"Ошибка: Не удалось получить системную информацию :("},
			})

			log.Println("failed to get message payload")
			return
		}

		switch payload.Command {
		case "start":
			messages.Start(a, object)
			return
		case "features":
			messages.Features(a, object)
			return
		case "functions":
			messages.Functions(a, object)
			return
		case "developer":
			messages.Developer(a, object)
			return
		case "geo":
			messages.Geo(a, object)
			return
		case "random_photo_cats":
			messages.RandomPhotoCats(a, object)
			return
		case "random_quote":
			messages.RandomQuote(a, object)
			return
		case "previous":
			messages.Previous(a, object)
			return
		default:
			_, _ = a.MessagesSend(url.Values{
				"peer_id": {fmt.Sprint(object.Message.PeerId)},
				"message": {"Ошибка: Не удалось обработать кнопку :("},
			})

			log.Println("invalid message payload")
			return
		}
	}

	switch object.Message.Text {
	case "Начать":
		messages.Start(a, object)
		return
	default:
		messages.InvalidCommand(a, object)
		return
	}
}
