package tasks

import (
	"backend/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type TasksHandler struct {
	service *services.TasksService
}

func NewTasksHandler(service *services.TasksService) *TasksHandler {
	return &TasksHandler{service: service}
}

func (h *TasksHandler) Tasks(c *gin.Context) {
	t, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching tasks"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *TasksHandler) TasksById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	t, err := h.service.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no task found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching task"})
		}
		return
	}
	c.JSON(http.StatusOK, t)
}
