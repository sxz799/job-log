package cron

import (
	"github.com/robfig/cron"
	"job-log/model/entity"
	"job-log/service"
)

func InitCron() {
	c := cron.New()
	clipboardService := service.ClipboardService{}
	_ = c.AddFunc("0 50 23 * *", func() {
		_ = clipboardService.Add(entity.Clipboard{})
	})
	c.Start()
}
