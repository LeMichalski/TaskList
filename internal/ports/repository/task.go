package repository

import "github.com/LeMichalski/TaskList/internal/domains"

type TaskRepository interface{
	Create(domains.Task)

}

type taskRepository struct{
	 
}

func NewTaskRepository() TaskRepository{
 	return taskRepository{}

}

func (t taskRepository) Create(domains.Task){

}