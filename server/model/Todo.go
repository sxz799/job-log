package model

type Todo struct {
	Id      int    `json:"id" gorm:"autoIncrement"`
	Message string `json:"message"`
}
