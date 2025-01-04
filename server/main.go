package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"job-log/config"
	"job-log/cron"
	"job-log/router"
	"job-log/util"
	"log"
)

//go:embed dist
var content embed.FS

func main() {
	log.Println("正在连接并初始化数据库...")
	util.InitDB()
	log.Println("添加定时任务...")
	cron.InitCron()

	gin.SetMode(config.GinMode)
	r := gin.Default()
	r.Use(cors.Default())
	router.RegRouter(r)
	router.RegWebRouter(r, content)
	log.Println("服务启动中,当前使用端口：", config.ServerPort)
	err := r.RunTLS(":"+config.ServerPort, "cert/pem.pem", "cert/key.key")
	if err != nil {
		log.Println("未找到证书文件，以http运行！")
		_ = r.Run(":" + config.ServerPort)
	}
}
