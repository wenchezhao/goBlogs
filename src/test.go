// @Title test.go
// @Description:
// @Author  zwc
// @Date 2021/7/20
// @Update  zwc  2021/7/20

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
