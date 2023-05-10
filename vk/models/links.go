package models

type Link struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Photo       Photo  `json:"photo"`
	IsFavorite  bool   `json:"is_favorite"`
}
