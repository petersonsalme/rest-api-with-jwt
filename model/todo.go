package model

// Todo represents a task to do
type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
