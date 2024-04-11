package controller

import (
	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get all invoices
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get single invoice
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create invoice
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Update invoice
	}
}
