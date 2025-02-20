package tests

import (
	Service "databaseService/pkg/service"
	"databaseService/pkg/utilities"
	"databaseService/tests/suite"
	"testing"
)

func Test_CreateTask_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &Service.CreateTaskRequest{
		Title:       "Задача 1",
		Description: "Описание задачи 1",
	}

	response, err := st.ClientAPI.CreateTask(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_CreateTask_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := new(Service.CreateTaskRequest)

	response, err := st.ClientAPI.CreateTask(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_GetTasks_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	response, err := st.ClientAPI.GetTasks(ctx, nil)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_GetTask_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &Service.GetTaskRequest{
		Id: 5,
	}

	response, err := st.ClientAPI.GetTask(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_UpdateTask_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &Service.UpdateTaskRequest{
		Id:          5,
		Title:       "Обновленная запись 1",
		Description: "Обновленное описание 1",
		Status:      "in_progress",
	}

	response, err := st.ClientAPI.UpdateTask(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_DeleteTask_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &Service.DeleteTaskRequest{
		Id: 4,
	}

	response, err := st.ClientAPI.DeleteTask(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}
