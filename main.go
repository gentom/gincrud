package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"hello": "Hi Gin-Gonic",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})

	router.Run(":8080")
}
