package keyboard

import (
	"encoding/json"
	"vkIntership/vk/helpers"
)

func New(oneTime, inLine bool) Keyboard {
	if oneTime && inLine {
		oneTime = false
	}

	return Keyboard{
		OneTime: oneTime,
		Inline:  inLine,
	}
}

func (k *Keyboard) AddButton(button Button) error {
	if len(k.Buttons) >= 6 {
		return helpers.ErrorKeyboardContainTooMuchRows
	}

	var count int

	for _, row := range k.Buttons {
		count += len(row)
	}

	if count >= 10 {
		return helpers.ErrorKeyboardContainsTooMuchButtons
	}

	if len(k.Buttons) < 1 {
		k.Buttons = make([][]Button, 1)
	}

	lastElement := len(k.Buttons) - 1
	if len(k.Buttons[lastElement]) < 6 {
		k.Buttons[lastElement] = append(k.Buttons[lastElement], button)
	} else {
		k.Buttons = append(k.Buttons, []Button{button})
	}

	return nil
}

func (k *Keyboard) AddRow() error {
	if len(k.Buttons) >= 6 {
		return helpers.ErrorKeyboardContainTooMuchRows
	}

	k.Buttons = append(k.Buttons, []Button{})
	return nil
}

func (k *Keyboard) ToJSON() string {
	for i, row := range k.Buttons {
		for j, button := range row {
			var payload string
			if button.Action.Payload != nil {
				jsonBytes, _ := json.Marshal(button.Action.Payload)
				payload = string(jsonBytes)
			}

			k.Buttons[i][j].Action.Payload = payload
		}
	}

	jsonBytes, _ := json.Marshal(k)
	return string(jsonBytes)
}
