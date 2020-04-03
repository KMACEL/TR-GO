package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login gelen değeri işleyen yapı
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// Örnek post sorgusunun Body kısmına yazılacak JSON  : {"user": "mert", "password": "123"}
	//Content-Type application/json olması gerek
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		// Gelen değer "user" yada "password" içermezse geriye açıklayıcı bir hata döndürüyor.
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gelen değeri kontrol ediyor, eğer kullanıcı adı ve şifre yanlışsa "unauthorized" hatası döndürüyor
		if json.User != "mert" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		//  Her şey yolundaysa "you are logged in" değerini döndürüyor
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Örnek post sorgusunun Body kısmına yazılacak XML :
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>mert</user>
	//		<password>123</password>
	//	</root>
	// Content-Type application/xml yada text/xml olması gerek
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "mert" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Örnek post sorgusunun Body kısmına yazılacak HTML form : user=mert&password=123
	// Content-Type header olması gerek
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "mert" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}
