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
		g.POST("/update", api.Update)
		g.GET("/finish/:id", api.Finish)
		g.GET("/unfinish/:id", api.UnFinish)
		g.GET("/:id", api.Get)
		g.DELETE("/:id", api.Del)

	}
}
