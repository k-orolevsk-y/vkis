package messages

import (
	"fmt"
	"net/url"
	"vkIntership/core/keyboards"
	"vkIntership/quotes"
	"vkIntership/vk/api"
	"vkIntership/vk/longpoll"
)

func RandomQuote(a api.Api, object *longpoll.MessageNewObject) {
	quote, err := quotes.Get()
	if err != nil {
		_, _ = a.MessagesSend(url.Values{
			"peer_id":  {fmt.Sprint(object.Message.PeerId)},
			"message":  {"🔴 Мне не удалось найти случайную цитату :("},
			"keyboard": {keyboards.RemoveMessage()},
		})
		return
	}

	var author string
	if quote.QuoteAuthor != "" {
		author = fmt.Sprintf("\n\n— %v", quote.QuoteAuthor)
	}

	_, _ = a.MessagesSend(url.Values{
		"peer_id":  {fmt.Sprint(object.Message.PeerId)},
		"message":  {fmt.Sprintf("🟢 %v%v", quote.QuoteText, author)},
		"keyboard": {keyboards.RandomQuote(quote.QuoteLink)},
	})

	return
}
