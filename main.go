package main

import (
	"fmt"
	"os"

	"github.com/breakinferno/golog/config"
	"github.com/breakinferno/golog/db"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	// 初始化配置
	config.Init()

	// 初始化数据库
	dbClient, err := db.Init()
	if err != nil {
		fmt.Printf("[Main] Init DB error: %+v", err)
		os.Exit(1)
	}
	defer dbClient.Close()
	// 初始化api服务
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}
