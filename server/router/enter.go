package router

import (
	"github.com/gin-gonic/gin"
	"job-log/router/Clipboard"
	"job-log/router/Todo"
)

func RegRouter(e *gin.Engine) {
	Todo.Todo(e)
	Clipboard.Clipboard(e)
}
