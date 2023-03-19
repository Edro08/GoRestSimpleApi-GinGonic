package create

import (
	Task "GoRestSimpleApi/internal/task"
	"context"
)

//New Task
type TaskCreateService struct {
	taskRepository Task.TaskSaveRepository
}

func NewTaskService(TaskRepository Task.TaskSaveRepository) TaskCreateService {
	return TaskCreateService{
		taskRepository: TaskRepository,
	}
}

func (s TaskCreateService) Save(ctx context.Context, Id, Name, Content string) error {
	task, err := Task.NewTask(Id, Name, Content)

	if err != nil {
		return err
	}

	return s.taskRepository.Save(ctx, task)
}
