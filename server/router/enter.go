package router

import (
	"github.com/gin-gonic/gin"
	"todo-demo/router/Todo"
)

func RegRouter(e *gin.Engine) {
	Todo.Todo(e)
}
