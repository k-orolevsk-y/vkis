package keyboards

import "vkIntership/vk/keyboard"

func RandomQuote(link string) string {
	k := keyboard.New(false, true)

	_ = k.AddButton(
		keyboard.NewButtonOpenLink("Цитата на сайте", link, nil),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonCallback("💨 Удали сообщение", Payload{Command: "remove_message"}, keyboard.NegativeColor),
	)

	return k.ToJSON()
}
