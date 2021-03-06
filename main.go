package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kholes/gin-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}