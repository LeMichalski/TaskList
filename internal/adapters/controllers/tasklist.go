package controller

import (
	"github.com/LeMichalski/TaskList/internal/domains"
	"github.com/LeMichalski/TaskList/internal/ports/usecase"
)

type TaskController interface{
	Create(domains.Task)

}

type taskController struct{
	taskUseCase usecase.TaskUseCase 
}

func NewTaskController() TaskUseCase{
 	return taskController{}

}

func (t taskController) Create(task domains.Task){ 
	t.taskUseCase.Create(task)
}