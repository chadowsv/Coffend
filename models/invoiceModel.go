package models

import (
	"time"
)

type Invoice struct {
	InvoiceID        int       `json:"invoice_id"`
	OrderID          int       `json:"order_id"`
	Iva              bool      `json:"iva"`
	Total            float64   `json:"total"`
	PaymentMethod    *string   `json:"paymentMethod" validate:"eq=TRANS|eq=CASH|eq="`
	PaymentStatus    bool      `json:"paymentStatus" validate:"required"`
	Payment_due_date time.Time `json:"payment_due_date"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at"`
}
