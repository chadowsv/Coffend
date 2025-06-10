package models

import "time"

type Table struct {
	TableID      int       `json:"table_id"`
	NumberGuests *int      `json:"number_guests" validate:"required"`
	Status       bool      `json:"status" validate:"required"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}
