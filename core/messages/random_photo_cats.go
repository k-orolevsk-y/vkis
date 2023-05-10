package messages

import (
	"fmt"
	"net/url"
	"os"
	"vkIntership/cats"
	"vkIntership/core/keyboards"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func RandomPhotoCats(a api.Api, object *longpoll.MessageNewObject) {
	m, _ := a.MessagesSendWithPeerIds(url.Values{
		"peer_ids": {fmt.Sprint(object.Message.PeerId)},
		"message":  {"🧬 Ищу лучших котиков в этом мире..."},
		"keyboard": {keyboards.Functions()},
	})

	errorUploadFunc := func(fileName string) {
		if fileName != "" {
			_ = os.Remove(fileName)
		}

		_, _ = a.MessagesEdit(url.Values{
			"peer_id":                 {fmt.Sprint(object.Message.PeerId)},
			"conversation_message_id": {fmt.Sprint(m.Response[0].ConversationMessageId)},
			"message":                 {"⚠️ Не удалось отправить картинку котика :("},
			"keyboard":                {keyboards.Functions()},
		})
	}

	fileName, err := cats.Download()
	if err != nil {
		errorUploadFunc(fileName)
		return
	}

	server, err := a.PhotosGetMessagesUploadServer(url.Values{
		"peer_id": {fmt.Sprint(object.Message.PeerId)},
	})
	if err != nil {
		errorUploadFunc(fileName)
		return
	}

	fileUpload, err := os.Open(fileName)
	if err != nil {
		errorUploadFunc(fileName)
		return
	}

	result, err := a.PhotosUpload(server, fileUpload)
	if err != nil {
		errorUploadFunc(fileName)
		return
	}

	photo, err := a.PhotosSaveMessagesPhoto(url.Values{
		"photo":  {result.Photo},
		"server": {fmt.Sprint(result.Server)},
		"hash":   {result.Hash},
	})
	if err != nil || len(photo.Response) < 1 {
		errorUploadFunc(fileName)
		return
	}
	_ = os.Remove(fileName)

	_, _ = a.MessagesEdit(url.Values{
		"peer_id":                 {fmt.Sprint(object.Message.PeerId)},
		"conversation_message_id": {fmt.Sprint(m.Response[0].ConversationMessageId)},
		"message":                 {"🐈 Держи котика."},
		"attachment":              {photo.Response[0].ToAttachment()},
		"keyboard":                {keyboards.RemoveMessage()},
	})

	return
}
