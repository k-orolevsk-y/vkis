package events

import "encoding/json"

func NewShowSnackbar(text string) Event {
	return Event{
		Type: ShowSnackbarType,
		Text: text,
	}
}

func NewOpenLink(link string) Event {
	return Event{
		Type: OpenLinkType,
		Link: link,
	}
}

func NewOpenApp(appId, ownerId int, hash string) Event {
	return Event{
		Type:    OpenAppType,
		AppId:   appId,
		OwnerId: ownerId,
		Hash:    hash,
	}
}

func (e *Event) ToJSON() string {
	jsonBytes, _ := json.Marshal(e)
	return string(jsonBytes)
}
