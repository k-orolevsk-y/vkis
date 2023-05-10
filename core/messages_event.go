package core

import (
	"fmt"
	"log"
	"net/url"
	"vkIntership/core/callback"
	"vkIntership/core/keyboards"
	"vkIntership/vk/events"
	"vkIntership/vk/longpoll"
)

func (c Core) messagesEvent(object *longpoll.MessageEventObject) {
	a := c.bot.GetApi()

	var payload keyboards.Payload
	if err := object.GetPayload(&payload); err != nil {
		evt := events.NewShowSnackbar("Ошибка: не удалось получить информацию о кнопке.")
		_, _ = a.MessagesSendMessageEventAnswer(url.Values{
			"event_id":   {object.EventId},
			"peer_id":    {fmt.Sprint(object.PeerId)},
			"user_id":    {fmt.Sprint(object.UserId)},
			"event_data": {evt.ToJSON()},
		})

		log.Println("failed to get command payload")
		return
	}

	switch payload.Command {
	case "remove_message":
		callback.RemoveMessage(a, object)
		return
	default:
		evt := events.NewShowSnackbar("Ошибка: неизвестная кнопка.")
		_, _ = a.MessagesSendMessageEventAnswer(url.Values{
			"event_id":   {object.EventId},
			"peer_id":    {fmt.Sprint(object.PeerId)},
			"user_id":    {fmt.Sprint(object.UserId)},
			"event_data": {evt.ToJSON()},
		})

		log.Println("invalid command payload")
		return
	}
}
