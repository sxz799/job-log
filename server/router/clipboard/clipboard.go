package clipboard

import (
	"github.com/gin-gonic/gin"
	"job-log/api/clipboard"
)

func Clipboard(e *gin.Engine) {
	g := e.Group("/api/clipboard")
	{
		g.GET("/list", clipboard.List)
		g.POST("", clipboard.Add)
		g.PUT("", clipboard.Update)
		g.GET("", clipboard.Get)
		g.DELETE("/:id", clipboard.Del)
	}
}
