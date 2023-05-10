package messages

import (
	"fmt"
	"log"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func Start(a api.Api, object *longpoll.MessageNewObject) {
	users, err := a.UsersGet([]string{fmt.Sprint(object.Message.FromId)}, "")
	if err != nil || len(users.Response) < 1 {
		_, _ = a.MessagesSend(url.Values{
			"peer_id": {fmt.Sprint(object.Message.PeerId)},
			"message": {"Ошибка: Не удалось получить о тебе информацию :("},
		})

		log.Printf("%v - error get user info", err)
		return
	}

	user := users.Response[0]
	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {fmt.Sprintf("Привет %v %v!", user.FirstName, user.LastName)},
		"keyboard": {keyboards.Start()},
	})

	return
}
