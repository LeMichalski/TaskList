package usecase

import (
	"github.com/LeMichalski/TaskList/internal/domains"
	"github.com/LeMichalski/TaskList/internal/ports/repository"
)

type TaskUseCase interface{
	Create(domains.Task)

}

type taskUseCase struct{
	taskRepository repository.TaskRepository 
}

func NewTaskUseCase() TaskUseCase{
 	return taskUseCase{}

}

func (t taskUseCase) Create(task domains.Task){ 
	t.taskRepository.Create(task)
}