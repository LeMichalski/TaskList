package main

import (
	"github.com/LeMichalski/TaskList/internal/adapters/controllers"
	"github.com/LeMichalski/TaskList/internal/domains"
	"github.com/gin-gonic/gin"
)

func main() {
	controller := controllers.NewTaskController()
	r := gin.Default()
	r.POST("/task", func(c *gin.Context) {
		controller.Create(domains.Task{})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
