package ginapp

import (
	"main/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	WorkerApp interfaces.WorkerApp
	Router    *gin.Engine
}

func (a *Api) Init() {
	router := gin.Default()

	router.GET("/tasks", a.GetTasks)
	router.POST("/tasks", a.StartTask)
	router.DELETE("/tasks", a.StopTask)

	a.Router = router
}

func (a *Api) GetTasks(c *gin.Context) {
	tasks := a.WorkerApp.GetTasks()
	c.JSON(http.StatusOK, gin.H{"method": "GET", "tasks": tasks})
}

func (a *Api) StartTask(c *gin.Context) {
	a.WorkerApp.StartTask()
	c.JSON(http.StatusOK, gin.H{"method": "POST", "message": "Task added"})
}

func (a *Api) StopTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}
