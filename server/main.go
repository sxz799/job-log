package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"job-log/config"
	"job-log/router"
	"job-log/service"
	"job-log/util"
	"log"
	"os"
)

func main() {
	log.Println("正在连接数据库...")
	util.InitDB()
	log.Println("正在检查表结构...")
	util.InitDBTables()

	gin.SetMode(config.GinMode)
	r := gin.Default()

	_, err := os.Stat("dist")
	if err == nil {
		r.LoadHTMLGlob("dist/index.html")
		r.Static("/dist", "dist")
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
		log.Println("已开启前后端整合模式！")
	}

	r.Use(cors.Default())

	router.RegRouter(r)

	c := cron.New()
	clipboardService := service.ClipboardService{}
	_ = c.AddFunc("0 50 23 * *", func() {
		_ = clipboardService.Add()
	})
	c.Start()

	log.Println("定时任务启动成功,服务启动成功,当前使用端口：", config.ServerPort)
	err = r.RunTLS(":"+config.ServerPort, "cert/pem.pem", "cert/key.key")
	if err != nil {
		log.Println("未找到证书文件，以http运行！")
		_ = r.Run(":" + config.ServerPort)
	}
}
