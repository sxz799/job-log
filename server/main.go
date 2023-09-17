package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"job-log/config"
	"job-log/router"
	"job-log/util"
	"log"
)

func main() {
	log.Println("正在连接数据库...")
	util.InitDB()
	log.Println("正在检查表结构...")
	util.InitDBTables()

	gin.SetMode(config.GinMode)
	r := gin.Default()

	r.Use(cors.Default())

	router.RegRouter(r)

	//TODO: 定时任务 晚上自动添加记录
	//c := cron.New()
	//c.AddFunc("0 50 23 * *", model.AutoTagImportant)
	//c.Start()

	err := r.Run(":" + config.ServerPort)
	if err != nil {
		return
	}
}
