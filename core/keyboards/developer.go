package keyboards

import "vkIntership/vk/keyboard"

func Developer() string {
	k := keyboard.New(false, true)

	_ = k.AddButton(keyboard.NewButtonOpenLink("💢 Telegram", "https://korolevsky.me/tg", nil))
	_ = k.AddRow()
	_ = k.AddButton(keyboard.NewButtonOpenLink("⚙️ Github", "https://korolevsky.me/gh", nil))

	return k.ToJSON()
}
