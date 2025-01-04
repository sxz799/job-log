package entity

type Clipboard struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type ClipboardLog struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
	DeleteAt string `json:"delete_at" gorm:"index"`
}
