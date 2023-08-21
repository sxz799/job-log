package service

import (
	"time"
	"todo-demo/model"
	"todo-demo/util"
)

type TodoService struct {
}

func (ts *TodoService) List() (todos []model.Todo, err error) {
	err = util.DB.Where("delete_at=?", "").Order("status").Order("id DESC").Find(&todos).Error
	return
}

func (ts *TodoService) Add(t model.Todo) (err error) {
	t.Status = "N"
	t.CreateAt = time.Now().Format(time.DateTime)
	err = util.DB.Create(&t).Error
	return
}

func (ts *TodoService) Update(t model.Todo) (err error) {
	t.UpdateAt = time.Now().Format(time.DateTime)
	err = util.DB.UpdateColumns(&t).Error
	return
}

func (ts *TodoService) Del(id int) (err error) {
	todo := model.Todo{
		Id:       id,
		DeleteAt: time.Now().Format(time.DateTime),
	}
	err = util.DB.UpdateColumns(&todo).Error
	return
}

func (ts *TodoService) Get(id int) (todo model.Todo, err error) {
	todo.Id = id
	err = util.DB.First(&todo).Error
	return
}
