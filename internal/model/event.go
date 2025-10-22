package model

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}