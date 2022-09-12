package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	driftRouter := gin.Default()
	driftRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := driftRouter.Run(":5500")
	if err != nil {
		log.Fatal(err)
	}
}
