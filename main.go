package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run("localhost:5000"); err != nil {
		log.Panicf(err.Error())
	}
}
