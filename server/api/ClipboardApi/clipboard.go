package ClipboardApi

import (
	"github.com/gin-gonic/gin"
	"job-log/model/common/request"
	"job-log/model/common/response"
	"job-log/model/entity"
	"job-log/service"
	"strconv"
)

var cs service.ClipboardService

func List(c *gin.Context) {
	var pi request.PageInfo
	err := c.ShouldBindQuery(&pi)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := cs.List(pi); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

func Update(c *gin.Context) {
	var clip entity.Clipboard
	err := c.ShouldBind(&clip)

	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	cs.Update(clip)
	response.OkWithMessage("更新成功", c)
}

func Get(c *gin.Context) {
	clip := cs.Get()
	response.OkWithData(clip, c)
}
func Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = cs.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
