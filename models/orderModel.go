package models

import "time"

type Order struct {
	OrderID    int       `json:"order_id"`
	OrderDate  time.Time `json:"order_date"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	//UserID     int       `json:"user_id"`
	Total   float64 `json:"total"`
	TableID int     `json:"table_id" validate:"required"`
}
