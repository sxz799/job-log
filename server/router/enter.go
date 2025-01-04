package router

import (
	"github.com/gin-gonic/gin"
	"job-log/router/clipboard"
	"job-log/router/todo"
)

func RegRouter(e *gin.Engine) {
	todo.Todo(e)
	clipboard.Clipboard(e)
}
