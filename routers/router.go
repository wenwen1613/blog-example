package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/wenwen1613/blog-example/docs"
	"github.com/wenwen1613/blog-example/middleware/jwt"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)

	appv1 := r.Group("/api/v1")
	appv1.Use(jwt.JWT())
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
