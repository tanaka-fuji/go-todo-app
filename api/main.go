package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
