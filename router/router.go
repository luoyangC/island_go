package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"island/api"
	_ "island/docs"
	"island/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	{
		// 获取验证码
		v1.POST("code", api.SendCode)
		// 获取定场诗
		v1.GET("sentence", api.RandomSentence)

		// 用户登录
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// 获取用户列表
		v1.GET("users", api.UserList)
		// 获取指定用户详情
		v1.GET("user/:id", api.UserDetail)

		// 获取话题列表
		v1.GET("topics", api.TopicList)
		// 获取话题详情
		v1.GET("topic/:id", api.TopicDetail)

		// 获取文章列表
		v1.GET("articles", api.ArticleList)
		// 获取文章详情
		v1.GET("article/:id", api.ArticleDetail)
	}

	// 需要登录保护的
	v1.Use(middleware.JWTAuth())
	{
		// 获取当前用户信息
		v1.GET("user", api.UserInfo)
		// 更新当前用户信息
		v1.PUT("user", api.UserUpdate)
		// 删除当前用户
		v1.DELETE("user", api.UserDelete)

		// 新建话题
		v1.POST("topic", api.TopicCreate)
		// 更新话题
		v1.PUT("topic/:id", api.TopicUpdate)
		// 删除话题
		v1.DELETE("topic/:id", api.TopicDelete)

		// 新建文章
		v1.POST("article", api.ArticleCreate)
		// 更新文章
		v1.PUT("article/:id", api.ArticleUpdate)
		// 删除文章
		v1.DELETE("article/:id", api.ArticleDelete)
	}
	return r
}
