package core

import (
	"log"
	"vkIntership/vk"
	"vkIntership/vk/longpoll"
)

func New(accessToken, version string) (Core, error) {
	bot, err := vk.NewBot(accessToken, version)
	if err != nil {
		return Core{}, err
	}

	return Core{
		bot: bot,
	}, nil
}

func (c Core) Run() error {
	lp := c.bot.GetLongPoll()

	cfg, err := lp.NewUpdateConfig()
	if err != nil {
		return err
	}
	channelUpdates := lp.GetUpdatesChan(cfg)

	for update := range channelUpdates {
		switch update.Type {
		case longpoll.MessageNew:
			object, err := update.GetMessageNewObject()
			if err != nil {
				log.Printf("failed to get new message object")
			}

			go c.messagesNew(&object)
			break
		case longpoll.MessageEvent:
			object, err := update.GetMessageEventObject()
			if err != nil {
				log.Printf("failed to get event message object")
			}

			go c.messagesEvent(&object)
			break
		}
	}

	return nil
}
