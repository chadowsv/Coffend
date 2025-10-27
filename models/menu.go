package models

import "time"

type Menu struct {
	MenuID     int       `gorm:"primaryKey;autoIncrement" json:"menu_id"`
	Name       string    `gorm:"type:nvarchar(100);not null" json:"name"`
	MenuStatus bool      `gorm:"not null" json:"menu_status"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`

	Foods []Food `gorm:"foreignKey:MenuID" json:"foods,omitempty"`
}
