package routes

import (
	"github.com/restaurent_table_booking/middlewares" // Import middleware

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.AdminOnly) // Gắn middleware vào nhóm router admin

	adminGroup.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Admin Dashboard!"})
	})
	// adminRoutes.GET("/dashboard", AdminDashboardHandler)
}
