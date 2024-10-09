package controllers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
	"strconv"
)

var tasks = []models.Task{}
var taskCategories = []models.TaskCategory{}
var currentTaskID = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {
	for i, task := range tasks {
		taskCategories := getCategoriesForTask(task.ID)
		tasks[i].Categories = taskCategories
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getCategoriesForTask(taskID int) []models.Category {
	var categories []models.Category

	for _, taskCategory := range taskCategories {
		if taskCategory.TaskID == taskID {
			for _, category := range categories {
				if category.ID == taskCategory.CategoryID {
					categories = append(categories, category)
				}
			}
		}
	}

	return categories
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if task.Task == "" || task.UserID == 0 {
		http.Error(w, "Task and UserID cannot be empty", http.StatusBadRequest)
		return
	}

	task.ID = currentTaskID
	currentTaskID++
	tasks = append(tasks, task)

	for _, category := range task.Categories {
		taskCategory := models.TaskCategory{
			TaskID:     task.ID,
			CategoryID: category.ID,
		}
		taskCategories = append(taskCategories, taskCategory)
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Task == "" {
				http.Error(w, "Task cannot be empty", http.StatusBadRequest)
				return
			}

			tasks[i] = updatedTask

			taskCategories = removeCategoriesForTask(id)
			for _, category := range updatedTask.Categories {
				taskCategory := models.TaskCategory{
					TaskID:     id,
					CategoryID: category.ID,
				}
				taskCategories = append(taskCategories, taskCategory)
			}

			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func removeCategoriesForTask(taskID int) []models.TaskCategory {
	var newTaskCategories []models.TaskCategory

	for _, taskCategory := range taskCategories {
		if taskCategory.TaskID != taskID {
			newTaskCategories = append(newTaskCategories, taskCategory)
		}
	}

	return newTaskCategories
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskCategories = removeCategoriesForTask(id)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
