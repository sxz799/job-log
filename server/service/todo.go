package service

import (
	"time"
	"todo-demo/model/common/request"
	"todo-demo/model/common/response"
	"todo-demo/model/entity"
	"todo-demo/util"
)

type TodoService struct {
}

func (ts *TodoService) List(pi request.PageInfo) (response.PageResult, error) {
	var todos []entity.Todo
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.Page - 1)
	db := util.DB.Model(&entity.Todo{})
	db = db.Where("delete_at=?", "")
	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Order("status").Order("id DESC").Find(&todos).Error
	return response.PageResult{
		List:     todos,
		Total:    total,
		Page:     pi.Page,
		PageSize: pi.PageSize}, err
}

func (ts *TodoService) Add(t entity.Todo) (err error) {
	t.Status = "N"
	t.CreateAt = time.Now().Format(time.DateTime)
	err = util.DB.Create(&t).Error
	return
}

func (ts *TodoService) Update(t entity.Todo) (err error) {
	t.UpdateAt = time.Now().Format(time.DateTime)
	err = util.DB.UpdateColumns(&t).Error
	return
}

func (ts *TodoService) Del(id int) (err error) {
	todo := entity.Todo{
		Id:       id,
		DeleteAt: time.Now().Format(time.DateTime),
	}
	err = util.DB.UpdateColumns(&todo).Error
	return
}

func (ts *TodoService) Get(id int) (todo entity.Todo, err error) {
	todo.Id = id
	err = util.DB.First(&todo).Error
	return
}
