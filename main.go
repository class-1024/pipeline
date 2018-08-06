package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/index", GetHome)
	router.Run(":80")
}
func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": c.DefaultQuery("name", "welcome to kinot")})
}
