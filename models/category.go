package models

type Category struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type TaskCategory struct {
	TaskID     int `json:"task_id"`
	CategoryID int `json:"category_id"`
}
