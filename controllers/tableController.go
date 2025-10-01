package controllers

import (
	"coffend/database"
	"coffend/models"
	"context"
	"log"
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
		insertQuery := "INSERT INTO Tables (table_id, number_guests,status,created_at, updated_at) VALUES (@p1,@p2,@p3,@p4,@p5)"
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

		// Obtener ID de la URL
		tableID := c.Param("table_id")

		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Validación
		if validationErr := validate.Struct(table); validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		// Verificar que la mesa existe (usando el ID de la URL)
		var existingTableID int
		err := database.DB.QueryRowContext(ctx,
			"SELECT table_id FROM Tables WHERE table_id = @p1",
			tableID).Scan(&existingTableID)

		if err != nil {
			c.JSON(404, gin.H{"error": "Table not found"})
			return
		}

		// Actualizar campos
		table.Updated_at = time.Now()

		// Query de actualización
		updateQuery := `
            UPDATE Tables
            SET number_guests = @p1, 
                status = @p2, 
                updated_at = @p3
            WHERE table_id = @p4
        `

		// Ejecutar actualización
		result, err := database.DB.ExecContext(ctx, updateQuery,
			table.NumberGuests,
			table.Status,
			table.Updated_at,
			tableID, // Usar el ID de la URL, no del JSON
		)

		if err != nil {
			log.Printf("Error updating table: %v", err) // Log detallado
			c.JSON(500, gin.H{
				"error":   "Failed to update table",
				"details": err.Error(), // Solo para desarrollo
			})
			return
		}

		// Verificar si realmente se actualizó
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(404, gin.H{"error": "No changes made - table not found"})
			return
		}

		c.JSON(200, gin.H{
			"message":  "Table updated successfully",
			"table_id": tableID,
			"changes": gin.H{
				"number_guests": table.NumberGuests,
				"status":        table.Status,
			},
		})
	}
}
