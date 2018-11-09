package v1

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/gavinzhou/hello-gin/models"
	"github.com/gavinzhou/hello-gin/pkg/e"
	"github.com/gavinzhou/hello-gin/pkg/setting"
	"github.com/gavinzhou/hello-gin/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.Config.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
