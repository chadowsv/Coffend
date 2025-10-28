package models

import "time"

type Food struct {
	FoodID      int       `gorm:"primaryKey;autoIncrement" json:"food_id"`
	Name        string    `gorm:"type:nvarchar(100);not null" json:"name"`
	Description string    `gorm:"type:nvarchar(255)" json:"description"`
	Price       float64   `gorm:"type:money;not null" json:"price"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
	MenuID      *int      `json:"menu_id"`

	Menu       *Menu       `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:FoodID" json:"-"`
}
