package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
