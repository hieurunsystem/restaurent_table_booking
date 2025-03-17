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

	// 		ADMIN ROUTES
	// Users (Tất cả người dùng: Admin, Owner, Customer)
	server.GET("/admin/users", GetUsers)         // Danh sách người dùng
	server.GET("/admin/users/:user_id", GetUser) // Thông tin người dùng

	server.GET("/admin/owners", GetOwners)          // Danh sách chủ nhà hàng
	server.GET("/admin/owners/:owner_id", GetOwner) // Thông tin chủ nhà hàng

	server.GET("/admin/customers", GetCustomers)             // Danh sách khách hàng
	server.GET("/admin/customers/:customer_id", GetCustomer) // Thông tin khách hàng

	server.GET("/admin/restaurants", GetRestaurants)               // Danh sách nhà hàng
	server.GET("/admin/restaurants/:restaurant_id", GetRestaurant) // Thông tin nhà hàng

	server.GET("/admin/tables", GetTables)          // Danh sách bàn ăn
	server.GET("/admin/tables/:table_id", GetTable) // Thông tin bàn ăn

	// 		OWNER ROUTES
	// Restaurant Routes (Nhà hàng thuộc về một chủ sở hữu cụ thể)
	server.GET("/owners/:owner_id/restaurants", GetAllRestaurants)                  // Get all restaurants of an owner
	server.GET("/owners/:owner_id/restaurants/:restaurant_id", GetRestaurantByID)   // Get restaurant by ID of an owner
	server.POST("/owners/:owner_id/restaurants", CreateRestaurant)                  // Create a new restaurant for an owner
	server.PUT("/owners/:owner_id/restaurants/:restaurant_id", EditRestaurant)      // Edit a restaurant of an owner
	server.DELETE("/owners/:owner_id/restaurants/:restaurant_id", DeleteRestaurant) // Delete a restaurant of an owner

	// Table Routes (Bàn thuộc về nhà hàng cụ thể)
	server.GET("/restaurants/:restaurant_id/tables", GetAllTables)             // Get all tables in a restaurant
	server.GET("/restaurants/:restaurant_id/tables/:table_id", GetTableByID)   // Get table by ID in a restaurant
	server.POST("/restaurants/:restaurant_id/tables", CreateTable)             // Create new table in a restaurant
	server.PUT("/restaurants/:restaurant_id/tables/:table_id", EditTable)      // Edit table in a restaurant
	server.DELETE("/restaurants/:restaurant_id/tables/:table_id", DeleteTable) // Delete table in a restaurant

	// 		CUSTOMER ROUTES
	server.POST("/restaurants/:restaurant_id/bookings", CreateBooking)      // API đặt bàn
	server.GET("/customers/:customer_id/bookings", GetBookingsByCustomerID) // Lịch sử theo ID khách hàng

}
