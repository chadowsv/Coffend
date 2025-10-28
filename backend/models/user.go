package models

import (
	"time"
)

type User struct {
	UserID       int       `gorm:"primaryKey;autoIncrement" json:"user_id"`
	FirstName    string    `gorm:"type:nvarchar(100);not null" json:"first_name"`
	LastName     string    `gorm:"type:nvarchar(100)" json:"last_name"`
	Password     string    `gorm:"type:nvarchar(255);not null" json:"password"`
	Email        string    `gorm:"type:nvarchar(25)" json:"email"`
	Role         string    `gorm:"type:nvarchar(25);not null" json:"role"`
	Phone        string    `gorm:"type:nvarchar(10)" json:"phone"`
	Token        string    `gorm:"type:nvarchar(max)" json:"token"`
	RefreshToken string    `gorm:"type:nvarchar(max)" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
}
