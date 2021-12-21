package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Gin框架",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
