package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
			"ids":     ids,
			"names":   names,
		})
	})
	router.Run(":8080")
}
