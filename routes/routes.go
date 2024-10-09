package routes

import (
	"github.com/gorilla/mux"
	"go-crud/controllers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users", controllers.DeleteUser).Methods("DELETE")
}

func RegisterTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", controllers.DeleteTask).Methods("DELETE")
}

func RegisterCategoryRoutes(router *mux.Router) {
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories", controllers.DeleteCategory).Methods("DELETE")
}
