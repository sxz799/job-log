package service

import (
	"job-log/model/common/request"
	"job-log/model/common/response"
	"job-log/model/entity"
	"job-log/util"
	"time"
)

var currClip = entity.Clipboard{}

type ClipboardService struct {
}

func (ts *ClipboardService) List(pi request.PageInfo) (response.PageResult, error) {
	var clips []entity.ClipboardLog
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.Page - 1)
	db := util.DB.Model(&entity.ClipboardLog{})
	db = db.Where("delete_at=?", "")
	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Order("id DESC").Find(&clips).Error
	return response.PageResult{
		List:     clips,
		Total:    total,
		Page:     pi.Page,
		PageSize: pi.PageSize}, err
}

func (ts *ClipboardService) Add(c entity.Clipboard) (err error) {
	var clipLog = entity.ClipboardLog{}
	clipLog.Content = currClip.Content
	if c.Content != "" {
		clipLog.Content = c.Content
	}
	clipLog.CreateAt = time.Now().Format(time.DateTime)
	err = util.DB.Debug().Where(entity.ClipboardLog{Content: clipLog.Content}).FirstOrCreate(&clipLog).Error
	return
}

func (ts *ClipboardService) Del(id int) (err error) {
	clip := entity.ClipboardLog{
		Id: id,
	}
	err = util.DB.Delete(&clip).Error
	return
}

func (ts *ClipboardService) Get() (clip entity.Clipboard) {
	return currClip
}

func (ts *ClipboardService) Update(t entity.Clipboard) {
	currClip = t
	return
}
