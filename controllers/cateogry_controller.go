package controllers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
	"strconv"
)

var categories = []models.Category{}
var currentCategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	if category.Category == "" {
		http.Error(w, "Category cannot be empty", http.StatusBadRequest)
		return
	}

	category.ID = currentCategoryID
	currentCategoryID++
	categories = append(categories, category)
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var updatedCategory models.Category
	json.NewDecoder(r.Body).Decode(&updatedCategory)

	for i, category := range categories {
		if category.ID == id {
			if updatedCategory.Category == "" {
				http.Error(w, "Category cannot be empty", http.StatusBadRequest)
				return
			}
			categories[i] = updatedCategory
			json.NewEncoder(w).Encode(updatedCategory)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
