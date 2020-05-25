package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenwen1613/blog-example/pkg/setting"
	"github.com/wenwen1613/blog-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	appv1 := r.Group("/api/v1")
	{
		appv1.GET("/tags", v1.GetTags)
		appv1.POST("/tags", v1.AddTag)
		appv1.DELETE("/tags", v1.DeleteTags)
		appv1.PUT("/tags", v1.EditTags)
	}
	return r
}
