package domain

import "time"

type Task struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	ParentID  *int       `json:"parent_id,omitempty"` // Using pointer to allow null value
	Status    bool       `json:"status"`              // Default false (0)
	IsDelete  bool       `json:"is_delete"`           // Default false (0)
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
