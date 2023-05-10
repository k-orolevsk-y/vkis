package models

import (
	"vkIntership/vk/helpers"
)

type Group struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	IsClosed   int    `json:"is_closed"`
	Type       string `json:"type"`
	Photo50    string `json:"photo_50"`
	Photo100   string `json:"photo_100"`
	Photo200   string `json:"photo_200"`
}

type ResultGroupsGetById struct {
	Error    helpers.Error `json:"error"`
	Response struct {
		Groups   []Group       `json:"groups"`
		Profiles []interface{} `json:"profiles"`
	} `json:"response"`
}

type ResultGetLongPollServer struct {
	Error    helpers.Error `json:"error"`
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}
