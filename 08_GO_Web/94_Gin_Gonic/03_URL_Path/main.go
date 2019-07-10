package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Bu sorgu "/user/mert" yapısını yakalar. Fakat "/user/"" veya "/user" denmemesi gerekir
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Merhaba %s", name)
	})

	// "/user/mert/" yapısından sonra gelecek parametreler için kullanılır. Örneğin : /user/mert/send
	// Fakat "/user/mert" denirse, bir önceki sorgu işleme girecektir.
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " işlemi " + action
		c.String(http.StatusOK, message)
	})

	// "/user/x/y" formatında olan her URL yolu için geçerlidir. "/user/:name/*action" gibi kontrolller ile kullanılabilir
	router.POST("/user/:name/*action", func(c *gin.Context) {
		c.String(http.StatusOK, strconv.FormatBool(c.FullPath() == "/user/:name/*action"))
	})

	router.Run(":8080")
}
