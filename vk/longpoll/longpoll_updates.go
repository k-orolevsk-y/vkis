package longpoll

import (
	"encoding/json"
	"vkIntership/vk/helpers"
	"vkIntership/vk/models"
)

type MessageNewObject struct {
	Message    models.Message `json:"message"`
	ClientInfo ClientInfo     `json:"client_info"`
}

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangId         int      `json:"lang_id"`
}

func (u Update) GetMessageNewObject() (MessageNewObject, error) {
	if u.Type != MessageNew {
		return MessageNewObject{}, helpers.ErrorInvalidObject
	}

	var object MessageNewObject
	if err := json.Unmarshal(u.Object, &object); err != nil {
		return MessageNewObject{}, err
	}

	return object, nil
}

type MessageEventObject struct {
	UserId                int             `json:"user_id"`
	PeerId                int             `json:"peer_id"`
	EventId               string          `json:"event_id"`
	Payload               json.RawMessage `json:"payload"`
	ConversationMessageId int             `json:"conversation_message_id"`
}

func (u Update) GetMessageEventObject() (MessageEventObject, error) {
	if u.Type != MessageEvent {
		return MessageEventObject{}, helpers.ErrorInvalidObject
	}

	var object MessageEventObject
	if err := json.Unmarshal(u.Object, &object); err != nil {
		return MessageEventObject{}, err
	}

	return object, nil
}

func (e MessageEventObject) GetPayload(object any) error {
	return helpers.GetPayload(e.Payload, &object)
}
