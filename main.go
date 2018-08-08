package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/index", GetHome)
	mqtt := router.Group("/mqtt/api/")
	{
		mqtt.POST("info", GetMqtt)
	}
	iot := router.Group("/iot/api/")
	{
		iot.POST("info", GetIot)
	}
	router.Run(":80")
}
func GetMqtt(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": c.DefaultQuery("name", "wehcome to mqtt api")})
}
func GetIot(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": c.DefaultQuery("name", "wehcome to iot api")})
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": c.DefaultQuery("name", "welcome to kinot index")})
}
