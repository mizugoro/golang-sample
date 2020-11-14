package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//htmlのディレクトリを指定
	r.LoadHTMLFiles("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			//htmlに渡す変数を定義
			"message": "hello gin",
		})
	})
	r.Run(":3030")
}
