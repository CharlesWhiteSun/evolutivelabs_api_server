package routers

import (
	"api_server/src/ginApiServer/pkg/setting"
	v1 "api_server/src/ginApiServer/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/sentences", v1.GetSentence)
		apiv1.GET("/sentences/:param", v1.GetItsThisForThatSentence)
	}

	return r
}
