package keyboards

import "vkIntership/vk/keyboard"

func Start() string {
	k := keyboard.New(false, false)

	_ = k.AddButton(
		keyboard.NewButtonText("Что ты умеешь?", Payload{Command: "features", Other: nil}, keyboard.PrimaryColor),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonText("Список функций", Payload{Command: "functions", Other: nil}, keyboard.SecondaryColor),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonText("Разработчик", Payload{Command: "developer", Other: nil}, keyboard.PositiveColor),
	)
	_ = k.AddRow()
	_ = k.AddButton(
		keyboard.NewButtonLocation(Payload{Command: "geo"}),
	)

	return k.ToJSON()
}
