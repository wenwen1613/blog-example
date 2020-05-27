package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wenwen1613/blog-example/pkg/setting"
	"github.com/wenwen1613/blog-example/routers/api"
	v1 "github.com/wenwen1613/blog-example/routers/api/v1"
)

//InitRouter 初始化路由规则
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	appv1 := r.Group("/api/v1")
	{
		appv1.GET("/tags", v1.GetTags)
		appv1.POST("/tags", v1.AddTag)
		appv1.DELETE("/tags", v1.DeleteTags)
		appv1.PUT("/tags", v1.EditTags)

		appv1.GET("/article", v1.GetArticle)
		appv1.GET("/articles", v1.GetArticles)
		appv1.POST("/article", v1.AddArticle)
		appv1.PUT("/article", v1.EditArticle)
		appv1.DELETE("/article", v1.DeleteArticle)
	}
	return r
}
