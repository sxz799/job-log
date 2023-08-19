package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"todo-demo/config"
	"todo-demo/router"
	"todo-demo/util"
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

	err := r.Run(":" + config.ServerPort)
	if err != nil {
		return
	}
}
