package repository

import (
	"database/sql"
	"errors"
	"to-do-list-app/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Insert Task
func (r *TaskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	query := `INSERT INTO tasks (title, parent_id, user_id, status, is_delete, created_at, updated_at)
	          VALUES (?, ?, ?, 0, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, title, parent_id, user_id, status, is_delete, created_at, updated_at`

	var createdTask domain.Task
	err := r.db.QueryRow(query, task.Title, task.ParentID, task.UserID).
		Scan(&createdTask.ID, &createdTask.Title, &createdTask.ParentID, &createdTask.UserID,
			&createdTask.Status, &createdTask.IsDelete, &createdTask.CreatedAt, &createdTask.UpdatedAt)

	if err != nil {
		return domain.Task{}, err
	}

	return createdTask, nil
}

// Get Task by ID
func (r *TaskRepository) GetTaskByID(id int) (*domain.Task, error) {
	query := `SELECT id, title, parent_id, status, is_delete, user_id, created_at, updated_at, deleted_at 
              FROM tasks WHERE id = ?`

	row := r.db.QueryRow(query, id)
	var task domain.Task
	err := row.Scan(&task.ID, &task.Title, &task.ParentID, &task.Status, &task.IsDelete, &task.UserID, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

// Update Task
func (r *TaskRepository) UpdateTask(task domain.Task) error {
	query := `UPDATE tasks SET title = ?, parent_id = ?, status = ?, is_delete = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, task.Title, task.ParentID, task.Status, task.IsDelete, task.ID)
	return err
}

// Delete Task (soft delete)
func (r *TaskRepository) SoftDeleteTask(id int) error {
	_, err := r.db.Exec("UPDATE tasks SET is_delete = 1 WHERE id = ?", id)
	return err
}

// Get All Tasks
func (r *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	query := `SELECT id, title, COALESCE(parent_id, NULL), user_id, status, is_delete, created_at, updated_at 
			  FROM tasks WHERE is_delete = 0`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID, &task.Title, &task.ParentID, &task.UserID,
			&task.Status, &task.IsDelete, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
