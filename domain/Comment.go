package domain

import "time"

// Comment domain
type Comment struct {
	Text          string    `json:"text" validate:"required"`
	User          string    `json:"user_id" validate:"required"`
	Establishment string    `json:"establishment_id" validate:"required"`
	Purchase      string    `json:"purchase_id" validate:"required"`
	Score         int       `json:"score" validate:"required,max=5,min=1"`
	Created       time.Time `json:"created"`
	Deleted       bool      `json:"deleted"`
}
