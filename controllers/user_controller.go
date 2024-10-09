package controllers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
	"strconv"
)

var users = []models.User{}
var currentUserID = 1

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
		return
	}

	user.ID = currentUserID
	currentUserID++
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	for i, user := range users {
		if user.ID == id {
			if updatedUser.Username == "" || updatedUser.Password == "" {
				http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
				return
			}
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
