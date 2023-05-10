package vk

import (
	"vkIntership/vk/models"
	"vkIntership/vk/requests"
)

type Bot struct {
	Group     models.Group
	requester *requests.Requester

	Debug bool
}
