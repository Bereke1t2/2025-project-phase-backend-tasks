package router


import (
	"task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetUp_Router(router *gin.Engine) {
	taskController := controllers.NewTaskController()

	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks/:id", taskController.GetTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)
	router.GET("/tasks", taskController.ListTasks)
}
