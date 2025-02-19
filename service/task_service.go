package service

import (
	"errors"
	"to-do-list-app/domain"
	"to-do-list-app/dto"
	"to-do-list-app/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

// Create Task
func (s *TaskService) CreateTask(req dto.TaskCreateRequest) (dto.TaskCreateResponse, error) {
	task := domain.Task{
		Title:    req.Title,
		ParentID: req.ParentID,
		Status:   false,
		IsDelete: false,
		UserID:   req.UserID,
	}

	createdTask, err := s.taskRepo.CreateTask(task)
	if err != nil {
		return dto.TaskCreateResponse{}, err
	}

	// Konversi ke DTO Response
	response := dto.TaskCreateResponse{
		ID:        createdTask.ID,
		Title:     createdTask.Title,
		ParentID:  createdTask.ParentID,
		UserID:    createdTask.UserID,
		Status:    createdTask.Status,
		IsDelete:  createdTask.IsDelete,
		CreatedAt: createdTask.CreatedAt,
		UpdatedAt: createdTask.UpdatedAt,
	}

	return response, nil
}

// Get Task by ID
func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.taskRepo.GetTaskByID(id)
}

// Update Task
func (s *TaskService) UpdateTask(id int, req dto.TaskUpdateRequest) error {
	task, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task not found")
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.ParentID != nil {
		task.ParentID = req.ParentID
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.IsDelete != nil {
		task.IsDelete = *req.IsDelete
	}

	return s.taskRepo.UpdateTask(*task)
}

// Soft Delete Task
func (s *TaskService) SoftDeleteTask(id int) error {
	return s.taskRepo.SoftDeleteTask(id)
}

// Get All Tasks
func (s *TaskService) GetAllTasks() ([]dto.TaskResponse, error) {
	tasks, err := s.taskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			ParentID:  task.ParentID, // Pastikan ParentID ikut dimasukkan
			UserID:    task.UserID,
			Status:    task.Status,
			IsDelete:  task.IsDelete,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}

	return taskResponses, nil
}
