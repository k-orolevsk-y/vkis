package keyboard

func NewButtonText(label string, payload any, color ButtonColor) Button {
	return Button{
		Action: ButtonAction{
			Type:    TextType,
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}
}

func NewButtonOpenLink(label, link string, payload any) Button {
	return Button{
		Action: ButtonAction{
			Type:    OpenLinkType,
			Label:   label,
			Link:    link,
			Payload: payload,
		},
	}
}

func NewButtonLocation(payload any) Button {
	return Button{
		Action: ButtonAction{
			Type:    LocationType,
			Payload: payload,
		},
	}
}

func NewButtonVkPay(hash string, payload any) Button {
	return Button{
		Action: ButtonAction{
			Type:    VkPayType,
			Hash:    hash,
			Payload: payload,
		},
	}
}

func NewButtonOpenApp(label string, appId int, ownerId int, hash string, payload any) Button {
	return Button{
		Action: ButtonAction{
			Type:    OpenAppType,
			Label:   label,
			AppId:   appId,
			OwnerId: ownerId,
			Hash:    hash,
			Payload: payload,
		},
	}
}

func NewButtonCallback(label string, payload any, color ButtonColor) Button {
	return Button{
		Action: ButtonAction{
			Type:    CallBackType,
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}
}
