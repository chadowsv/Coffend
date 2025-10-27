package models

import "time"

type Order struct {
	OrderID   int       `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderDate time.Time `gorm:"not null" json:"order_date"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	TableID   *int      `json:"table_id"` // puede ser null
	Total     float64   `gorm:"type:money" json:"total"`

	Table      *Table      `gorm:"foreignKey:TableID" json:"table,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	Invoice    *Invoice    `gorm:"foreignKey:OrderID" json:"invoice,omitempty"`
}
