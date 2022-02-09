package middleware

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		sub := "admin"
		e := model.Casbin()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success := e.Enforce(sub, obj, act)
		if success {
			log.Println("恭喜您,权限验证通过")
			ctx.Next()
		} else {
			log.Printf("e.Enforce err: %s", "很遗憾,权限验证没有通过")
			response.ToErrorResponse(errcode.UnauthorizedAuthFail)
			ctx.Abort()
			return
		}
	}
}
