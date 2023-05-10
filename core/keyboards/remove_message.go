package keyboards

import "vkIntership/vk/keyboard"

func RemoveMessage() string {
	k := keyboard.New(false, true)
	_ = k.AddButton(
		keyboard.NewButtonCallback("💨 Удали сообщение", Payload{Command: "remove_message"}, keyboard.NegativeColor),
	)
	return k.ToJSON()
}
