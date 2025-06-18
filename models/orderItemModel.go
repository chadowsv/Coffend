package models

import "time"

type OrderItem struct {
	OrderItemID int       `json:"order_item_id"`
	Quantity    float64   `json:"quantity" validate:"required"`
	UnitPrice   float64   `json:"unit_price" validate:"required"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	OrderID     int       `json:"order_id" validate:"required"`
	FoodID      int       `json:"food_id" validate:"required"`
}
type OrderItemExtended struct {
	OrderItemID int       `json:"order_item_id"`
	OrderID     int       `json:"order_id"`
	FoodID      int       `json:"food_id"`
	FoodName    string    `json:"food_name"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unit_price"`
	Subtotal    float64   `json:"subtotal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
