package service

import (
	"todo-demo/model"
	"todo-demo/util"
)

type TodoService struct {
}

func (ts *TodoService) List() (todos []model.Todo, err error) {
	err = util.DB.Order("id DESC").Find(&todos).Error
	return
}

func (ts *TodoService) Get(id int) (todo model.Todo, err error) {
	todo.Id = id
	err = util.DB.First(&todo).Error
	return
}

func (ts *TodoService) Del(id int) (err error) {
	err = util.DB.Delete(model.Todo{}, "id=? ", id).Error
	return
}

func (ts *TodoService) Add(t model.Todo) (err error) {
	err = util.DB.Create(&t).Error
	return
}
