package Clipboard

import (
	"github.com/gin-gonic/gin"
	"job-log/api/ClipboardApi"
)

func Clipboard(e *gin.Engine) {
	g := e.Group("/api/clipboard")
	{
		g.GET("/list", ClipboardApi.List)
		g.PUT("/", ClipboardApi.Update)
		g.GET("/", ClipboardApi.Get)
		g.DELETE("/:id", ClipboardApi.Del)

	}
}
