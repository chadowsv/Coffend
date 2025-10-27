package models

import "time"

type Invoice struct {
	InvoiceID      int             `gorm:"primaryKey;autoIncrement" json:"invoice_id"`
	OrderID        int             `gorm:"not null" json:"order_id"`
	IVA            bool            `gorm:"not null" json:"iva"`
	Total          float64         `gorm:"not null" json:"total"`
	PaymentMethod  string          `gorm:"type:nvarchar(50)" json:"payment_method"`
	PaymentStatus  bool            `gorm:"not null" json:"payment_status"`
	PaymentDueDate time.Time       `gorm:"not null" json:"payment_due_date"`
	CreatedAt      time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"not null" json:"updated_at"`
	Details        []InvoiceDetail `gorm:"foreignKey:InvoiceID" json:"details,omitempty"`

	TableNumber  int         `gorm:"-" json:"table_number,omitempty"`
	OrderDetails []OrderItem `gorm:"foreignKey:OrderID" json:"order_details,omitempty"`
}
