package routers

import (
	"github.com/gavinzhou/hello-gin/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter(config setting.AppConfig) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(config.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
