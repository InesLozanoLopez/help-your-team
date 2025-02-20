package tasks

import (
	"backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, s *services.TasksService) {
	h := NewTasksHandler(s)

	tr := r.Group("tasks")
	{
		tr.GET("/all", h.Tasks)
		tr.GET("/:id", h.TasksById)
	}
}
