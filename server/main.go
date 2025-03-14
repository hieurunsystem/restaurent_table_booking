package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/db"
	"github.com/restaurent_table_booking/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Chỉ cho phép frontend của bạn truy cập
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Cho phép gửi cookie qua CORS
	}))

	// Đăng ký các routes
	routes.Routes(server)      // Các route chung
	routes.AdminRoutes(server) // Các route yêu cầu quyền Admin

	server.Run()
}
