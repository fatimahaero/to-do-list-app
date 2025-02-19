package dto

import "time"

type TaskCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	ParentID *int   `json:"parent_id,omitempty"` // Bisa null
	UserID   int    `json:"user_id" validate:"required"`
}

type TaskCreateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ParentID  *int      `json:"parent_id,omitempty"` // Bisa null
	UserID    int       `json:"user_id"`
	Status    bool      `json:"status"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskUpdateRequest struct {
	Title    *string `json:"title,omitempty"`
	ParentID *int    `json:"parent_id,omitempty"`
	Status   *bool   `json:"status,omitempty"`
	IsDelete *bool   `json:"is_delete,omitempty"`
}

type TaskResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ParentID  *int      `json:"parent_id,omitempty"` // Bisa null
	UserID    int       `json:"user_id"`
	Status    bool      `json:"status"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
