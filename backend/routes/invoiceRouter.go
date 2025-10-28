package routes

import (
	controller "coffend/controllers"
	"coffend/middleware"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", middleware.AuthRequired(), middleware.AdminOnly(), controller.GetAllInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.GetInvoiceByID())
	incomingRoutes.POST("/invoices", middleware.AuthRequired(), middleware.AdminOnly(), controller.PostInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.PatchInvoiceByID())
	incomingRoutes.DELETE("/invoices/:invoice_id", middleware.AuthRequired(), middleware.AdminOnly(), controller.DeleteInvoiceByID())
}
