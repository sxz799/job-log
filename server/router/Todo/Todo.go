package Todo

import (
	"github.com/gin-gonic/gin"
	"job-log/api/TodoApi"
)

func Todo(e *gin.Engine) {
	g := e.Group("/api/todo")
	{
		g.GET("/list", TodoApi.List)
		g.POST("/", TodoApi.Add)
		g.PUT("/:id", TodoApi.Update)
		g.GET("/:id", TodoApi.Get)
		g.DELETE("/:id", TodoApi.Del)

	}
}
