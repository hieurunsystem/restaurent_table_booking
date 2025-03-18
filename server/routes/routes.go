package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/middlewares"
)

func Routes(server *gin.Engine) {
	// Generic
	server.GET("/", Home)

	// Authentication
	server.POST("/login", Login)
	server.POST("/register", Register)
	server.POST("/logout", Logout)
	server.GET("/me", middlewares.AuthMiddleware(), GetUserProfile)

	// Users routes (View - Add - Edit - Delete)

	AdminRoutes(server) // Các route yêu cầu quyền Admin
	OwnerRoutes(server)

	

}
