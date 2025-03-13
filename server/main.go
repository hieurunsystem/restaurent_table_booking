package main

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurent_table_booking/db"
	"github.com/restaurent_table_booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.Routes(server)
	server.Run()
}
