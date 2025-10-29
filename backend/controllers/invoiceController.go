package controllers

import (
	"coffend/backend/database"
	"coffend/backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoices []models.Invoice

		if err := database.DB.WithContext(ctx).Find(&invoices).Error; err != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}

		c.JSON(200, invoices)
	}
}

func GetInvoiceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		if err := database.DB.WithContext(ctx).
			Preload("Order.OrderItems.Food").
			Preload("Order.Table").
			First(&invoice, "invoice_id = ?", invoiceId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
			return
		}

		if invoice.Order != nil && invoice.Order.Table != nil && invoice.Order.Table.TableID != 0 {
			invoice.TableNumber = invoice.Order.Table.TableID
		}

		if invoice.Order != nil {
			invoice.OrderDetails = invoice.Order.OrderItems
		}

		c.JSON(http.StatusOK, invoice)
	}
}

func PostInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var newInvoice models.Invoice

		if err := c.BindJSON(&newInvoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var order models.Order
		if err := database.DB.WithContext(ctx).
			Preload("OrderItems").
			First(&order, "order_id = ?", newInvoice.OrderID).Error; err != nil {
			c.JSON(400, gin.H{"error": "Order not found"})
			return
		}

		if len(order.OrderItems) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Order has no items"})
			return
		}

		// Calculate total
		var total float64
		var iva float64 = 1.15
		for _, item := range order.OrderItems {
			total += float64(item.Quantity) * float64(item.UnitPrice)
		}

		if newInvoice.IVA {
			total *= iva
		}

		// Create the invoice
		invoice := models.Invoice{
			OrderID:        newInvoice.OrderID,
			IVA:            newInvoice.IVA,
			Total:          total,
			PaymentMethod:  newInvoice.PaymentMethod,
			PaymentStatus:  newInvoice.PaymentStatus,
			PaymentDueDate: newInvoice.PaymentDueDate,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := database.DB.WithContext(ctx).Create(&invoice).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
			return
		}

		if err := database.DB.WithContext(ctx).
			Model(&models.Order{}).
			Where("order_id = ?", newInvoice.OrderID).
			Update("total", total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order total"})
			return
		}

		if err := database.DB.WithContext(ctx).
			Preload("Order.OrderItems.Food").
			First(&invoice, invoice.InvoiceID).Error; err == nil {
			c.JSON(http.StatusCreated, invoice)
			return
		}

		c.JSON(http.StatusCreated, invoice)
	}
}

func PatchInvoiceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var inputInvoice models.Invoice
		var invoice models.Invoice

		if err := c.BindJSON(&inputInvoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(inputInvoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).First(&invoice, invoiceId).Error; err != nil {
			c.JSON(400, gin.H{"error": "Invoice not found"})
			return
		}

		if err := database.DB.WithContext(ctx).First(&models.Order{}, "order_id = ?", invoice.OrderID).Error; err != nil {
			c.JSON(400, gin.H{"error": "Order not found"})
			return
		}

		// Update invoice fields
		updates := map[string]interface{}{
			"order_id":         inputInvoice.OrderID,
			"payment_method":   inputInvoice.PaymentMethod,
			"payment_status":   inputInvoice.PaymentStatus,
			"payment_due_date": inputInvoice.PaymentDueDate,
			"updated_at":       time.Now(),
		}

		if err := database.DB.WithContext(ctx).Model(&invoice).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully", "invoice": invoice})
	}
}

func DeleteInvoiceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		if err := database.DB.WithContext(ctx).Delete(&invoice, invoiceId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
	}
}
