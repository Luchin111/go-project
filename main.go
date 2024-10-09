package main

import (
	"github.com/gorilla/mux"
	"go-crud/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Registrar rutas
	routes.RegisterUserRoutes(router)
	routes.RegisterTaskRoutes(router)
	routes.RegisterCategoryRoutes(router)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":8082", router))
}
