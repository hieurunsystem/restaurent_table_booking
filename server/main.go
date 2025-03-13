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
	server.Use(cors.Default())
	routes.Routes(server)
	server.Run()
}
