package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// MaxMultipartMemory (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Tek Dosya
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Belirtilen Bulunan konuma dosyayı kaydeder.
		c.SaveUploadedFile(file, file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! - %d", file.Filename, file.Size))

	})
	router.Run(":8080")
}
