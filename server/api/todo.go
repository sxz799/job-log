package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-demo/model"
	"todo-demo/service"
	"todo-demo/util"
)

var ts service.TodoService

func List(c *gin.Context) {
	if list, err := ts.List(); err == nil {
		util.ResponseOK(c, "查询成功！", list)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}
}
func Add(c *gin.Context) {
	var todo model.Todo
	err := c.ShouldBind(&todo)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = ts.Add(todo); err == nil {
		util.ResponseOK(c, "添加成功！", todo)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}

}
func Update(c *gin.Context) {
	var todo model.Todo
	err := c.ShouldBind(&todo)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = ts.Update(todo); err == nil {
		util.ResponseOK(c, "更新成功！", todo)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}
}
func Finish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = ts.Finish(id); err == nil {
		util.ResponseOK(c, "更新成功！", id)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if result, err := ts.Get(id); err == nil {
		util.ResponseOK(c, "查询成功！", result)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}

}
func Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = ts.Del(id); err == nil {
		util.ResponseOK(c, "删除成功！", nil)
	} else {
		util.ResponseError(c, http.StatusNotFound, err.Error())
	}

}
