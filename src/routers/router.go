package routers

import (
	_ "github.com/Wenchuan-Zhao/goBlogs/docs"
	"github.com/Wenchuan-Zhao/goBlogs/pkg/setting"
	"github.com/Wenchuan-Zhao/goBlogs/pkg/upload"
	"github.com/Wenchuan-Zhao/goBlogs/routers/api"
	v1 "github.com/Wenchuan-Zhao/goBlogs/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// InitRouter initialize routing informatio
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	// 图片路由
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// swagger 初始化和对应的路由规则
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 新增获取 token 的方法
	r.GET("/auth", api.GetAuth)
	// 增加上传照片的路由
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")

	// 中间件引用
	//apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
