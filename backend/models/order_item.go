package models

import "time"

type OrderItem struct {
	OrderItemID int       `gorm:"primaryKey;autoIncrement" json:"order_item_id"`
	Quantity    float64   `json:"quantity"`
	UnitPrice   float64   `gorm:"type:money" json:"unit_price"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
	OrderID     int       `gorm:"not null" json:"order_id"`
	FoodID      *int      `json:"food_id"`

	Order *Order `gorm:"foreignKey:OrderID" json:"-"`
	Food  *Food  `gorm:"foreignKey:FoodID" json:"food,omitempty"`
}
