package keyboard

type Keyboard struct {
	OneTime bool       `json:"one_time,omitempty"`
	Buttons [][]Button `json:"buttons"`
	Inline  bool       `json:"inline"`
}

type Button struct {
	Action ButtonAction `json:"action"`
	Color  ButtonColor  `json:"color,omitempty"`
}

type ButtonAction struct {
	Type    ButtonType `json:"type"`
	Label   string     `json:"label,omitempty"`
	Payload any        `json:"payload"`

	AppId   int    `json:"app_id,omitempty"`
	OwnerId int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Link    string `json:"link,omitempty"`
}

type ButtonType string
type ButtonColor string
