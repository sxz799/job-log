package Todo

import (
	"github.com/gin-gonic/gin"
	"todo-demo/api"
)

func Todo(e *gin.Engine) {
	g := e.Group("/api/todo")
	{
		g.GET("/list", api.List)
		g.POST("/add", api.Add)
		g.GET("/:id", api.Get)
		g.DELETE("/:id", api.Del)
	}
}
