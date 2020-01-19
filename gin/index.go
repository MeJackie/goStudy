package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 静态资源加载，本例为css,js以及资源图片
	router.StaticFS("/public", http.Dir("F:/me/GoJob/src/github.com/ffhelicopter/tmm/website/static"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:80
	router.Run(":80")
}