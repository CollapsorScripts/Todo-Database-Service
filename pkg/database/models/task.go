package models

import (
	"context"
	"databaseService/pkg/database"
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *Task) CreateTask(ctx context.Context) error {
	dbPool := database.GetPool()
	query := `
        INSERT INTO tasks (title, description, status, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, title, description, status,created_at, updated_at;
    `

	return dbPool.QueryRow(ctx, query, t.Title, t.Description, t.Status).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
}

func Tasks(ctx context.Context) ([]*Task, error) {
	dbPool := database.GetPool()

	query := `
        SELECT id, title, description, status, created_at, updated_at
        FROM tasks
        ORDER BY created_at DESC;
    `

	rows, err := dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := &Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func GetTask(ctx context.Context, ID int64) (*Task, error) {
	dbPool := database.GetPool()

	query := `
        SELECT id, title, description, status, created_at, updated_at
        FROM tasks
        WHERE id = $1;
    `

	task := new(Task)
	err := dbPool.QueryRow(ctx, query, ID).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return task, err
}

func (t *Task) UpdateTask(ctx context.Context) error {
	dbPool := database.GetPool()

	query := `
        UPDATE tasks
        SET title = $1,
            description = $2,
            status = $3,
            updated_at = NOW()
        WHERE id = $4
        RETURNING id, title, description, status, created_at, updated_at;
    `

	return dbPool.QueryRow(ctx, query, t.Title, t.Description, t.Status, t.ID).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
}

func DeleteTask(ctx context.Context, ID int64) {
	dbPool := database.GetPool()

	query := `
        DELETE FROM tasks
        WHERE id = $1;
    `

	dbPool.QueryRow(ctx, query, ID)
}
