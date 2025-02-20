package Service

import (
	"context"
	"databaseService/pkg/database/models"
	"databaseService/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *serverAPI) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	newTask := models.Task{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Status:      "new",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	if newTask.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Поле title не может быть пустым")
	}

	if newTask.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Поле description не может быть пустым")
	}

	if err := newTask.CreateTask(ctx); err != nil {
		logger.Info("Ошибка при создании модели: %v", err)
		return nil, status.Errorf(codes.Internal, "Ошибка при создании модели")
	}

	response := &CreateTaskResponse{
		Id:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      newTask.Status,
		CreatedAt:   newTask.CreatedAt.String(),
		UpdatedAt:   newTask.UpdatedAt.String(),
	}

	return response, nil
}

func (s *serverAPI) GetTasks(ctx context.Context, req *Empty) (*GetTasksResponse, error) {
	tasks, err := models.Tasks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Ошибка при поиске задач")
	}

	tasksResponse := make([]*Task, len(tasks))

	for i := 0; i < len(tasks); i++ {
		tasksResponse[i] = &Task{
			Id:          tasks[i].ID,
			Title:       tasks[i].Title,
			Description: tasks[i].Description,
			Status:      tasks[i].Status,
			CreatedAt:   tasks[i].CreatedAt.String(),
			UpdatedAt:   tasks[i].UpdatedAt.String(),
		}
	}

	response := &GetTasksResponse{
		Tasks: tasksResponse,
	}

	return response, nil
}

func (s *serverAPI) GetTask(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error) {
	task, err := models.GetTask(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	response := &GetTaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		UpdatedAt:   task.UpdatedAt.String(),
	}

	return response, nil
}

func (s *serverAPI) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	task := &models.Task{
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Status:      req.GetStatus(),
	}

	if err := task.UpdateTask(ctx); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Ошибка при обновлении задачи: %v", err)
	}

	response := &UpdateTaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		UpdatedAt:   task.UpdatedAt.String(),
	}

	return response, nil
}

func (s *serverAPI) DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	models.DeleteTask(ctx, req.GetId())

	response := &DeleteTaskResponse{
		Success: true,
	}

	return response, nil
}
