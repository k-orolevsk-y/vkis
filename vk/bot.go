package vk

import (
	"vkIntership/vk/api"
	"vkIntership/vk/helpers"
	"vkIntership/vk/longpoll"
	"vkIntership/vk/models"
	"vkIntership/vk/requests"
)

func NewBot(accessToken, version string) (*Bot, error) {
	return NewBotWithApiEndpoint(accessToken, version, helpers.APIEndpoint)
}

func NewBotWithApiEndpoint(accessToken, version, apiEndpoint string) (*Bot, error) {
	botApi := &Bot{}
	botApi.requester = requests.New(accessToken, apiEndpoint, version)
	botApi.requester.Debug = &botApi.Debug

	bot, err := botApi.getBot()
	if err != nil {
		return nil, err
	}

	botApi.Group = bot
	return botApi, nil
}

func (b *Bot) getBot() (models.Group, error) {
	a := b.GetApi()

	groups, err := a.GroupsGetById()
	if err != nil {
		return models.Group{}, err
	} else if len(groups.Response.Groups) < 1 {
		return models.Group{}, helpers.ErrorInvalidToken
	}

	return groups.Response.Groups[0], nil
}

func (b *Bot) GetApi() api.Api {
	return api.New(b.requester)
}

func (b *Bot) GetLongPoll() longpoll.LongPoll {
	a := b.GetApi()
	return longpoll.New(&a, b.requester, &b.Group)
}
