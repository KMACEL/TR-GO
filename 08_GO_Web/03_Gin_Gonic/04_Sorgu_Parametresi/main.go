package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Örnek url parametreleri :  localhost:8080/welcome?firstname=Mert&lastname=Acel
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Misafir") // "Misafir" default değer
		lastname := c.Query("lastname")                     // uzun hali  c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "Merhaba SN. %s %s, Hoş Geldiniz... ", firstname, lastname)
	})
	router.Run(":8080")
}
