package models

import "time"

type Table struct {
	TableID      int       `gorm:"primaryKey" json:"table_id"`
	NumberGuests int       `json:"number_guests"`
	Status       bool      `gorm:"not null" json:"status"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`

	Orders []Order `gorm:"foreignKey:TableID" json:"orders,omitempty"`
}
