package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 读取基本配置文件
	baseConfigPath := flag.String("bbb", "", "base config file path")
	deployConfigPath := flag.String("ccc", "", "deploy config file path which for overwriting base config")
	flag.Parse()
	fmt.Println(*baseConfigPath, *deployConfigPath)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
