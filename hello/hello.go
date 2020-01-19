package main

import (
	"fmt"
	"github.com/jian/study/hellolib"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("你好， 世界！ \n")
	fmt.Printf("2和3中最大的是 %d !",hellolib.Max(2,3))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
