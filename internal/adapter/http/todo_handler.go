package http

import (
	"errors"
	"feature-flag-poc/internal/application/port"
	"feature-flag-poc/internal/application/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	listTotoUsecase port.ListTodoUsecase
	startUpAt       time.Time
}

func NewTodoHandler(listTodoUsecase port.ListTodoUsecase) *todoHandler {
	return &todoHandler{
		listTotoUsecase: listTodoUsecase,
		startUpAt:       time.Now(),
	}
}

func (h *todoHandler) Register(r *gin.Engine) {
	r.GET("/todos", h.ListTodos)
	r.GET("/health", h.HealthCheck)
}

func (hdl *todoHandler) HealthCheck(c *gin.Context) {
	uptime := int64(time.Since(hdl.startUpAt).Seconds())
	resp := gin.H{
		"status": "Healthy",
		"uptime": uptime,
	}
	c.JSON(http.StatusOK, resp)
}

func (hdl *todoHandler) ListTodos(c *gin.Context) {
	todos, err := hdl.listTotoUsecase.Execute(c.Request.Context())
	if err != nil {
		if errors.Is(err, usecase.ErrFeatureIsDisabled) {
			c.JSON(http.StatusGone, gin.H{
				"code":    "GONE",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    todos,
		"message": "retrieve list todos successfully!",
	})
}
