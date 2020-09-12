package main

import (
	"github.com/GetStartWithGoLang/gin/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := handle.NewMember()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/member", func(c *gin.Context) {
		c.JSON(200, m.AllData())
	})
	r.Run() // listen on port 8080
}
