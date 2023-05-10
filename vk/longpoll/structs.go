package longpoll

import (
	"encoding/json"
	"vkIntership/vk/api"
	"vkIntership/vk/models"
	"vkIntership/vk/requests"
)

type LongPoll struct {
	shutdownChannel chan interface{}
	group           *models.Group
	api             *api.Api
	req             *requests.Requester
}

type UpdateConfig struct {
	Key    string
	Server string
	Ts     string
	Wait   int
}

type Update struct {
	Type    TypeEvent       `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupId int64           `json:"group_id"`
	EventId string          `json:"event_id"`
	Version string          `json:"v"`
}

type Updates struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
	Failed  int      `json:"failed"`
}

type TypeEvent string

type UpdatesChannel <-chan Update
