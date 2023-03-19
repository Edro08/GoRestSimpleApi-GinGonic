package update

import (
	"GoRestSimpleApi/internal/task"
	Task "GoRestSimpleApi/internal/task"
	"context"
)

//New Task
type TaskUpdateService struct {
	taskRepository Task.TaskUpdateRepository
}

func NewTaskService(TaskRepository Task.TaskUpdateRepository) TaskUpdateService {
	return TaskUpdateService{
		taskRepository: TaskRepository,
	}
}

func (s TaskUpdateService) Update(ctx context.Context, task task.Task) error {

	return s.taskRepository.Update(ctx, task)
}
