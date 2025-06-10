package models

import "time"

type Menu struct {
	MenuID     int       `json:"menu_id"`
	Name       string    `json:"name" validate:"required"`
	MenuStatus bool      `json:"menu_status" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
