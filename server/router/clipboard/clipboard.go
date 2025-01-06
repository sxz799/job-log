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
		g.GET("", clipboard.Get)
		g.DELETE("/:id", clipboard.Del)
		g.GET("/ws", clipboard.Ws)
	}
	//创建一个ws 实时更新信息

}
