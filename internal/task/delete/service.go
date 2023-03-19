package delete

import (
	Task "GoRestSimpleApi/internal/task"
	"context"
)

//New Task
type TaskDeleteService struct {
	taskRepository Task.TaskDeleteRepository
}

func NewTaskService(TaskRepository Task.TaskDeleteRepository) TaskDeleteService {
	return TaskDeleteService{
		taskRepository: TaskRepository,
	}
}

func (s TaskDeleteService) Delete(ctx context.Context, Id string) error {

	return s.taskRepository.Delete(ctx, Id)
}
