package models

import "time"

type Note struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"upated_at"`
}
