package todo

import (
	"github.com/gin-gonic/gin"
	"job-log/api/todo"
)

func Todo(e *gin.Engine) {
	g := e.Group("/api/todo")
	{
		g.GET("/list", todo.List)
		g.POST("", todo.Add)
		g.PUT("/:id", todo.Update)
		g.GET("/:id", todo.Get)
		g.DELETE("/:id", todo.Del)
	}
}
