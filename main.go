package main

import (
	"go-crud/database"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	router := gin.Default()

	routes.InitializeRoutes(router)

	router.Run(":8080")
}
