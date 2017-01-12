package main

import (
	jsonp "github.com/madglory/gin-jsonp"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.New()
	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(jsonp.Handler())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8088") // listen and server on 0.0.0.0:8080
}
