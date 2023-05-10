package models

import (
	"vkIntership/vk/helpers"
)

type Message struct {
	Attachments           []MessageAttachment `json:"attachments"`
	ConversationMessageID int                 `json:"conversation_message_id"`
	Date                  int                 `json:"date"`
	FromId                int                 `json:"from_id"`
	FwdMessages           []Message           `json:"fwd_Messages"`
	ReplyMessage          *Message            `json:"reply_message"`
	ID                    int                 `json:"id"`
	IsCropped             int                 `json:"is_cropped"`
	Payload               string              `json:"payload"`
	PeerId                int                 `json:"peer_id"`
	RandomId              int                 `json:"random_id"`
	Text                  string              `json:"text"`
	UpdateTime            int                 `json:"update_time"`
	Geo                   MessageGeo          `json:"geo"`
}

func (m Message) GetPayload(object any) error {
	return helpers.GetPayload([]byte(m.Payload), &object)
}

type MessageAttachment struct {
	Link    Link    `json:"link"`
	Photo   Photo   `json:"photo"`
	Sticker Sticker `json:"sticker"`
	Type    string  `json:"type"`
}

type MessageGeo struct {
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Place struct {
		City    string `json:"city"`
		Country string `json:"country"`
		Title   string `json:"title"`
	} `json:"place"`
	Type string `json:"type"`
}

func (g *MessageGeo) Has() bool {
	return g.Type != ""
}

type ResultMessagesSend struct {
	Error    helpers.Error `json:"error"`
	Response int           `json:"response"`
}

type ResultMessagesSendPeerIds struct {
	Error    helpers.Error `json:"error"`
	Response []struct {
		PeerId                int `json:"peer_id"`
		MessageId             int `json:"message_id"`
		ConversationMessageId int `json:"conversation_message_id"`
	} `json:"response"`
}

type ResultMessagesSendMessageEventAnswer struct {
	Error    helpers.Error `json:"error"`
	Response int           `json:"response"`
}

type ResultMessageEdit struct {
	Error    helpers.Error `json:"error"`
	Response int           `json:"response"`
}

type ResultMessageDelete struct {
	Response map[string]int `json:"response"`
}
