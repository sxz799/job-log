package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"job-log/router/clipboard"
	"job-log/router/todo"
	"log"
	"net/http"
)

func RegRouter(e *gin.Engine) {
	e.Use(cors())
	todo.Todo(e)
	clipboard.Clipboard(e)
}

func RegWebRouter(e *gin.Engine, content embed.FS) {
	temp := template.Must(template.New("").ParseFS(content, "dist/index.html"))
	e.SetHTMLTemplate(temp)
	distFS, _ := fs.Sub(content, "dist")
	e.StaticFS("/dist", http.FS(distFS))
	e.NoRoute(func(context *gin.Context) {
		context.HTML(200, "index.html", "")
	})
	log.Println("已开启前后端整合模式！")
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
