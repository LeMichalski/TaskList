package usecase

import (
	"github.com/LeMichalski/TaskList/internal/domains"
	"github.com/LeMichalski/TaskList/internal/ports/repository"
)

type TaskUseCase interface {
	Create(domains.Task)
}

type taskUseCaseImpl struct {
	taskRepositoryImpl repository.TaskRepository
}

func NewTaskUseCase() TaskUseCase {
	return taskUseCaseImpl{}

}

func (t taskUseCaseImpl) Create(task domains.Task) {
	t.taskRepositoryImpl.Create(task)
}
