package routers

import (
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.AccessLog())
	//r.Use(middleware.Recovery())

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))

	r.POST("/auth", api.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	casbin := v1.NewCasbin()
	test := v1.NewTest()
	messageTag := v1.NewMessageTag()
	message := v1.NewMessage()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	apiv1.Use(middleware.CasbinHandler())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

		// 测试路由
		apiv1.GET("/hello", test.Get)
		// 权限策略管理
		apiv1.POST("/casbin", casbin.Create)
		apiv1.POST("/casbin/list", casbin.List)
	}

	apiv2 := r.Group("/api/v1")
	{
		apiv2.POST("/messageTag/create", messageTag.Create)
		apiv2.POST("/message/create", message.Create)
		apiv2.POST("/message/receiveMessage", message.ReceiveMessage)

	}

	return r
}
