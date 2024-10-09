package controllers

import (
	"go-crud/database"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// Similar para Update y Delete

func UpdateUser(c *gin.Context) {
	var user models.User
	// Buscamos al usuario por su ID
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	// Actualizamos el usuario con los datos recibidos en la solicitud
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardamos los cambios en la base de datos
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser elimina un usuario por su ID
func DeleteUser(c *gin.Context) {
	var user models.User
	// Buscamos al usuario por su ID
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	// Eliminamos el usuario de la base de datos
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted!"})
}
