package models

type InvoiceDetail struct {
	InvoiceDetailID int        `gorm:"primaryKey;column:inovice_detail_id" json:"invoice_detail_id"`
	InvoiceID       int        `gorm:"column:invoice_id" json:"invoice_id"`
	ItemID          int        `gorm:"column:item_id" json:"item_id"`
	Quantity        int        `gorm:"column:quantity" json:"quantity"`
	Price           float64    `gorm:"type:decimal(10,2);column:price" json:"price"`
	Invoice         *Invoice   `gorm:"foreignKey:InvoiceID;references:InvoiceID" json:"invoice,omitempty"`
	OrderItem       *OrderItem `gorm:"foreignKey:ItemID;references:OrderItemID" json:"order_item,omitempty"`
}
