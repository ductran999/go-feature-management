package http

import "github.com/gin-gonic/gin"

func NewRouter(todoHandler *todoHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	todoHandler.Register(r)
	return r
}
