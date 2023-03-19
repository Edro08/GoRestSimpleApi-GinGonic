package read

import (
	"GoRestSimpleApi/internal/task"
	Task "GoRestSimpleApi/internal/task"
	"context"
)

//Task Get All
type TaskGetAllService struct {
	taskRepository Task.TaskGetAllRepository
}

func NewTaskService(TaskRepository Task.TaskGetAllRepository) TaskGetAllService {
	return TaskGetAllService{
		taskRepository: TaskRepository,
	}
}

func (s TaskGetAllService) ReadAll(ctx context.Context) ([]task.TaskDB, error) {

	return s.taskRepository.ReadAll(ctx)
}

//Task Get ID
type TaskGetIdService struct {
	taskRepository Task.TaskGetIdRepository
}

func NewTaskServiceID(TaskRepository Task.TaskGetIdRepository) TaskGetIdService {
	return TaskGetIdService{
		taskRepository: TaskRepository,
	}
}

func (s TaskGetIdService) ReadID(ctx context.Context, Id string) ([]task.TaskDB, error) {

	return s.taskRepository.ReadID(ctx, Id)
}
