package models

type Sticker struct {
	Images               []StickerImage `json:"images"`
	ImagesWithBackground []StickerImage `json:"images_with_background"`
	ProductId            int            `json:"product_id"`
	StickerId            int            `json:"sticker_id"`
	IsAllowed            bool           `json:"is_allowed"`
	AnimationUrl         string         `json:"animation_url"`
}

type StickerImage struct {
	Height float64 `json:"height"`
	Url    string  `json:"url"`
	Width  float64 `json:"width"`
	Type   string  `json:"type"`
}
