package model

type Todo struct {
	Id       int    `json:"id" gorm:"autoIncrement"`
	Status   string `json:"status"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at" gorm:"index"`
	FinishAt string `json:"finish_at"`
}
