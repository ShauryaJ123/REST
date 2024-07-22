package main

import (
	"abc.com/calc/db"
	"abc.com/calc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server:=gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}




