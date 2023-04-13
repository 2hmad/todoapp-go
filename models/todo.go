package models

type Todo struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Done  bool   `json:"done" db:"done"`

	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}
