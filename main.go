package main

import (
	"net/http"

	"github.com/gentom/gincrud/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		ctrl := task.TASK()

		tasks := ctrl.GetAllTasks()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"tasks": tasks,
		})
	})

	router.POST("/", func(c *gin.Context) {
		text := c.PostForm("text")

		ctrl := task.TASK()

		ctrl.AddTask(text)

		tasks := ctrl.GetAll()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"tasks": tasks,
		})
	})

	/* Router Practice
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"hello": "Hi Gin-Gonic",
		})
	})

	router.GET("/page", func(c *gin.Context) {
		c.String(200, "LOVE ELIXIR")
	})
	*/

	router.Run(":8080")
}
