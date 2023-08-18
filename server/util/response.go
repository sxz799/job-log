package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(context *gin.Context, httpStatus int, code int, success bool, msg string, data any) {
	context.JSON(httpStatus, gin.H{
		"code":    code,
		"success": success,
		"message": msg,
		"data":    data,
	})
}
func ResponseOK(context *gin.Context, msg string, data any) {
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"message": msg,
		"data":    data,
	})
}
func ResponseError(context *gin.Context, code int, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"success": false,
		"message": msg,
		"data":    nil,
	})
}
