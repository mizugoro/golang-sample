package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
	/*http.ListenAndServe(address,engin)を中身で呼んでいる*/
}

/*
net-http.go
html-template.goが参考になる
*/
