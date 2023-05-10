package models

import (
	"fmt"
	"vkIntership/vk/helpers"
)

type Photo struct {
	AlbumId   int         `json:"album_id"`
	Date      int         `json:"date"`
	Id        int         `json:"id"`
	OwnerId   int         `json:"owner_id"`
	Sizes     []PhotoSize `json:"sizes"`
	Text      string      `json:"text"`
	UserId    int         `json:"user_id"`
	HasTags   bool        `json:"has_tags"`
	AccessKey string      `json:"access_key"`
}

type PhotoSize struct {
	Height int    `json:"height"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

func (p Photo) ToAttachment() string {
	return fmt.Sprintf("photo%v_%v_%v", p.OwnerId, p.Id, p.AccessKey)
}

type ResultPhotosServer struct {
	Error    helpers.Error `json:"error"`
	Response struct {
		AlbumId   int    `json:"album_id"`
		UploadUrl string `json:"upload_url"`
		UserId    int    `json:"user_id"`
		GroupId   int    `json:"group_id"`
	} `json:"response"`
}

type ResultUploadServer struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

type ResultSaveMessagesPhoto struct {
	Response []Photo `json:"response"`
}
