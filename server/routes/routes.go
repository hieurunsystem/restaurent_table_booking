package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/middlewares"
)

func Routes(server *gin.Engine) {
	// User Management
	server.GET("/", Home)
	server.POST("/login", Login)
	server.POST("/register", Register)
	server.GET("/user_list", GetUser)
	server.POST("/logout", Logout)
	server.GET("/me", middlewares.AuthMiddleware(), GetUserProfile)

	// Restaurant Management
	server.GET("/restaurants/", GetAllRestaurants)
	server.POST("/restaurants/create", CreateRestaurant)
}
