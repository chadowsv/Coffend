package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InvoiceViewFormat struct {
	InvoiceID      int
	PaymentMethod  *string
	OrderID        int
	PaymentStatus  bool
	PaymentDue     float64
	TableNumber    int
	PaymentDueDate time.Time
	OrderDetails   []models.OrderItemExtended
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var invoices []models.Invoice
		queryInvoices := "SELECT * FROM Invoices"
		rows, err := database.DB.QueryContext(ctx, queryInvoices)
		if err != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var invoice models.Invoice
			err := rows.Scan(
				&invoice.InvoiceID,
				&invoice.OrderID,
				&invoice.Iva,
				&invoice.Total,
				&invoice.PaymentMethod,
				&invoice.PaymentStatus,
				&invoice.Payment_due_date,
				&invoice.Created_at,
				&invoice.Updated_at,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to parse invoice"})
				return
			}
			invoices = append(invoices, invoice)
		}
		c.JSON(200, invoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		invoiceId := c.Param("invoice_id")
		var invoice InvoiceViewFormat
		queryInvoice := "SELECT invoice_id, payment_method, order_id, payment_status, total FROM Invoices WHERE invoice_id = @p1"
		insertErr := database.DB.QueryRowContext(ctx, queryInvoice, invoiceId).Scan(
			&invoice.InvoiceID,
			&invoice.PaymentMethod,
			&invoice.OrderID,
			&invoice.PaymentStatus,
			&invoice.PaymentDue,
		)
		if insertErr != nil {
			c.JSON(404, gin.H{"error": "Invoice not found."})
			return
		}
		// 2. Obtener datos adicionales de la orden (por ejemplo, número de mesa)
		queryOrder := `SELECT table_id FROM Orders WHERE order_id = @p1`
		err := database.DB.QueryRowContext(ctx, queryOrder, invoice.OrderID).Scan(&invoice.TableNumber)
		if err != nil {
			c.JSON(500, gin.H{"error": "Order data not found."})
			return
		}

		// 3. Obtener detalles de los items de la orden usando tu función ItemsByOrder
		orderItems, err := ItemsByOrder(invoice.OrderID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching order items."})
			return
		}
		invoice.OrderDetails = orderItems

		// 4. Devolver la estructura completa como JSON
		c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var invoice models.Invoice

		// Validar entrada JSON
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var order models.Order
		erro := database.DB.QueryRowContext(ctx,
			"SELECT order_id FROM Orders WHERE order_id = @p1", invoice.OrderID).Scan(&order.OrderID)
		if erro != nil {
			c.JSON(400, gin.H{"error": "Order not found"})
			return
		}
		invoice.Created_at = time.Now()
		invoice.Updated_at = time.Now()
		insertQuery := "INSERT INTO Invoices (order_id,iva,total,payment_method,payment_status,payment_due_date,created_at,updated_at) VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8)"
		_, insertErr := database.DB.ExecContext(ctx, insertQuery,
			invoice.OrderID,
			invoice.Iva,
			invoice.Total,
			invoice.PaymentMethod,
			invoice.PaymentStatus,
			invoice.Payment_due_date,
			invoice.Created_at,
			invoice.Updated_at,
		)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect create invoice"})
			return
		}
		c.JSON(http.StatusCreated, invoice)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		// Validar entrada JSON
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Validar estructura
		if err := validate.Struct(invoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//Comprobar si la orden existe
		// Verificar que la orden existe
		var order models.Order
		erro := database.DB.QueryRowContext(ctx,
			"SELECT order_id FROM Orders WHERE order_id = @p1", invoice.OrderID).Scan(&order.OrderID)
		if erro != nil {
			c.JSON(400, gin.H{"error": "Order not found"})
			return
		}
		// Actualizar campos permitidos
		query := `
			UPDATE Invoices SET
				order_id = @p1,
				payment_method = @p2,
				payment_status = @p3,
				payment_due_date = @p4,
				updated_at = @p5
			WHERE invoice_id = @p6
		`

		_, err := database.DB.ExecContext(ctx, query,
			invoice.OrderID,
			invoice.PaymentMethod,
			invoice.PaymentStatus,
			invoice.Payment_due_date,
			time.Now(),
			invoiceId,
		)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update invoice"})
			return
		}

		c.JSON(200, gin.H{"message": "Invoice updated successfully"})
	}
}
