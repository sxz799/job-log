package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo-demo/model/common/request"
	"todo-demo/model/common/response"
	"todo-demo/model/entity"
	"todo-demo/service"
)

var ts service.TodoService

func List(c *gin.Context) {
	var pi request.PageInfo
	err := c.ShouldBindQuery(&pi)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := ts.List(pi); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}
func Add(c *gin.Context) {
	var todo entity.Todo
	err := c.ShouldBind(&todo)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ts.Add(todo); err == nil {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
func Update(c *gin.Context) {
	var todo entity.Todo
	err := c.ShouldBind(&todo)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ts.Update(todo); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := ts.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
func Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ts.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
