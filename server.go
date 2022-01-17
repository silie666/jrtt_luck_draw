package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("server/*")
	r.Static("/server", "./server")
	r.GET("/sign", func(c *gin.Context) {
		url := c.Query("url")
		c.HTML(http.StatusOK, "test.html", gin.H{
			"url": url,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
