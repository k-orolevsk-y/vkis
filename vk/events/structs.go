package events

type Event struct {
	Type    EventType `json:"type"`
	Text    string    `json:"text,omitempty"`
	Link    string    `json:"link,omitempty"`
	AppId   int       `json:"app_id,omitempty"`
	OwnerId int       `json:"owner_id,omitempty"`
	Hash    string    `json:"hash,omitempty"`
}

type EventType string
