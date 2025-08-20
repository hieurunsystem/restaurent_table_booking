package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/internal/controllers"
	"github.com/restaurent_table_booking/internal/utils/middlewares"
)

func Routes(server *gin.Engine) {
	// User Management
	server.GET("/", controllers.Home)
	server.POST("/login", controllers.LoginHandler)
	server.POST("/register", controllers.RegisterHandler)
	server.GET("/user_list", controllers.GetAllAccountsHandler)
	server.POST("/logout", controllers.LogoutHandler)
	server.GET("/me", middlewares.AuthMiddleware(), controllers.GetAllAccountsHandler)

	// Restaurant Management
	server.GET("/restaurants/", controllers.GetAllRestaurants)
	server.POST("/restaurants/create", controllers.CreateRestaurant)

	// Admin Management
	adminGroup := server.Group("/admin")
	adminGroup.Use(middlewares.AdminOnly) // Gắn middleware vào nhóm router admin

	adminGroup.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Admin Dashboard!"})
	})
}
