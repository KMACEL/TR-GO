package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getFonksyion",
	})
}

func postFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "postFonksyion",
	})
}

func putFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "putFonksyion",
	})
}

func deleteFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "deleteFonksyion",
	})
}

func patchFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "patchFonksyion",
	})
}

func headFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "headFonksyion",
	})
}

func optionsFonksyion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "optionsFonksyion",
	})
}

func noRouteFonksyion(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "noRouteFonksyion",
	})
}

func noMethodFonksyion(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "noMethodFonksyion",
	})
}

func main() {
	router := gin.Default()

	router.GET("/acel", getFonksyion)
	router.POST("/acel", postFonksyion)
	router.PUT("/acel", putFonksyion)
	router.DELETE("/acel", deleteFonksyion)
	router.PATCH("/acel", patchFonksyion)
	router.HEAD("/acel", headFonksyion)
	router.OPTIONS("/acel", optionsFonksyion)
	router.NoRoute(noRouteFonksyion)
	router.NoMethod(noMethodFonksyion)

	router.Run(":8080")
}
