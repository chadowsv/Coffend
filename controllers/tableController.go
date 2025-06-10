package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Contexto para que no dure tanto
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var tables []models.Table
		queryTables := "SELECT * FROM Tables"
		rows, err := database.DB.QueryContext(ctx, queryTables)
		if err != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var table models.Table
			err := rows.Scan(
				&table.TableID,
				&table.NumberGuests,
				&table.Status,
				&table.Created_at,
				&table.Updated_at,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to parse order"})
				return
			}
			tables = append(tables, table)
		}
		c.JSON(200, tables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := c.Param("table_id")
		var table models.Table
		queryTable := "	SELECT * FROM Tables WHERE table_id= @p1"
		err := database.DB.QueryRowContext(ctx, queryTable, tableId).Scan(
			&table.TableID,
			&table.NumberGuests,
			&table.Status,
			&table.Created_at,
			&table.Updated_at,
		)
		if err != nil {
			c.JSON(404, gin.H{"error": "Table not created."})
			return
		}
		c.JSON(200, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(table)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}
		table.Created_at = time.Now()
		table.Updated_at = time.Now()
		insertQuery := "INSERT INTO Tables (table_id, number_guests,status,created_at, updated_at) VALUES (@p1,@p2,@p3,@p4,@p5,"
		_, insertErr := database.DB.ExecContext(ctx, insertQuery,
			table.TableID,
			table.NumberGuests,
			table.Status,
			table.Created_at,
			table.Updated_at,
		)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect create table"})
			return
		}
		c.JSON(http.StatusCreated, table)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var table models.Table

		// Validación del formato JSON
		if err := c.BindJSON(&table); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Validación con validator
		validationErr := validate.Struct(table)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		tableID := c.Param("table_id")

		// Verificar que la mesa existe
		err := database.DB.QueryRowContext(ctx, "SELECT table_id FROM Tables WHERE table_id = @p1", table.TableID).Scan(&table.TableID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Table not found"})
			return
		}

		table.Updated_at = time.Now()

		// Query para actualizar la orden
		updateQuery := `
			UPDATE Tables
			SET table_id=@p1, number_guests=@p2, status=@p3, updated_at=@p4
			WHERE order_id=@p5
		`

		_, updateErr := database.DB.ExecContext(ctx, updateQuery,
			table.TableID,
			table.NumberGuests,
			table.Status,
			table.Updated_at,
			tableID,
		)

		if updateErr != nil {
			c.JSON(500, gin.H{"error": "Failed to update the table"})
			return
		}

		c.JSON(200, gin.H{"message": "Table updated successfully"})
	}
}
