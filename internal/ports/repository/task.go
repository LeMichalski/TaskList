package repository

import "github.com/LeMichalski/TaskList/internal/domains"

type TaskRepository interface {
	Create(domains.Task)
}

type taskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return taskRepositoryImpl{}

}

func (t taskRepositoryImpl) Create(domains.Task) {

}
