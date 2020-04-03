package main

import "github.com/gin-gonic/gin"

//ornek yapılacak
func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("/path/to/templates")

	r.Run(":8080")
}
