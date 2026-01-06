package http

import (
	"feature-flag-poc/internal/application/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	listTotoUsecase port.ListTodoUsecase
}

func NewTodoHandler(listTodoUsecase port.ListTodoUsecase) *todoHandler {
	return &todoHandler{
		listTotoUsecase: listTodoUsecase,
	}
}

func (h *todoHandler) Register(r *gin.Engine) {
	r.GET("/todos", h.ListTodos)
}

func (hdl *todoHandler) ListTodos(c *gin.Context) {
	todos, err := hdl.listTotoUsecase.Execute(c.Request.Context())
	if err != nil {
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
