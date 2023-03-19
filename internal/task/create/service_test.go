package create

import (
	Task "GoRestSimpleApi/internal/task"
	"GoRestSimpleApi/internal/task/platform/storage/storagemocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_TaskService_CreateTask_RepositoryError(t *testing.T) {
	taskId := "70"
	taskName := "Leer"
	taskContent := "Leer un libro cada dia"

	task, err := Task.NewTask(taskId, taskName, taskContent)
	require.NoError(t, err)

	taskRepositoryMock := new(storagemocks.TaskRepository)
	taskRepositoryMock.On("Save", mock.Anything, task).Return(errors.New("something unexpect happened"))
	taskService := NewTaskService(taskRepositoryMock)
	err = taskService.Save(context.Background(), taskId, taskName, taskContent)
	taskRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_TaskService_CreateTask_Succeed(t *testing.T) {
	taskId := "70"
	taskName := "Leer"
	taskContent := "Leer un libro cada dia"

	task, err := Task.NewTask(taskId, taskName, taskContent)
	require.NoError(t, err)

	taskRepositoryMock := new(storagemocks.TaskRepository)
	taskRepositoryMock.On("Save", mock.Anything, task).Return(nil)
	taskService := NewTaskService(taskRepositoryMock)
	err = taskService.Save(context.Background(), taskId, taskName, taskContent)
	taskRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
