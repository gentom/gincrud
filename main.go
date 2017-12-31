package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})

	router.Run(":8080")
}
