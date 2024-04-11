package controller

import (
	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get all tables
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get single table
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create table
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Update table
	}
}
