package models

type Sound struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Personality string `json:"personality"`
	CreatedAt   string `json:"created_at"`
}
