package models

type Task struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Task       string     `json:"task"`
	Categories []Category `json:"categories"` // Relación con categorías
}
