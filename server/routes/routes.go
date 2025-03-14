package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	// User Management
	server.GET("/", Home)
	server.POST("/login", Login)
	server.POST("/register", Register)
  server.GET("/user_list", GetUser)

	// Restaurant Management
	server.GET("/restaurants/", GetAllRestaurants)
	server.POST("/restaurants/create", CreateRestaurant)
}