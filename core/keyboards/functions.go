package keyboards

import "vkIntership/vk/keyboard"

func Functions() string {
	k := keyboard.New(false, false)

	_ = k.AddButton(
		keyboard.NewButtonText("Случайное фото котика", Payload{Command: "random_photo_cats"}, keyboard.PositiveColor),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonText("Случайная цитата", Payload{Command: "random_quote"}, keyboard.PositiveColor),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonText("Назад", Payload{Command: "previous"}, keyboard.SecondaryColor),
	)

	return k.ToJSON()
}
