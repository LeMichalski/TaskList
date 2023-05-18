package controllers

import (
	"net/http"

	"github.com/LeMichalski/TaskList/internal/domains"
	"github.com/LeMichalski/TaskList/internal/ports/usecase"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	Create(*gin.Context, domains.Task)
}

type taskControllerImpl struct {
	taskUseCaseImpl usecase.TaskUseCase
}

func NewTaskController() TaskController {
	return taskControllerImpl{}

}

func (t taskControllerImpl) Create(c *gin.Context, task domains.Task) {
	t.taskUseCaseImpl.Create(task)
	c.JSON(http.StatusCreated, nil)
}
