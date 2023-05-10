package callback

import (
	"fmt"
	"net/url"
	"vkIntership/vk/api"
	"vkIntership/vk/events"
	"vkIntership/vk/longpoll"
)

func RemoveMessage(a api.Api, object *longpoll.MessageEventObject) {
	_, _ = a.MessagesDelete(url.Values{
		"cmids":          {fmt.Sprint(object.ConversationMessageId)},
		"peer_id":        {fmt.Sprint(object.PeerId)},
		"delete_for_all": {"1"},
	})

	evt := events.NewShowSnackbar("💨 Сообщение удалено!")
	_, _ = a.MessagesSendMessageEventAnswer(url.Values{
		"event_id":   {object.EventId},
		"peer_id":    {fmt.Sprint(object.PeerId)},
		"user_id":    {fmt.Sprint(object.UserId)},
		"event_data": {evt.ToJSON()},
	})
}
