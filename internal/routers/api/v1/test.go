package v1

import (
	"blog-service/pkg/app"
	"github.com/gin-gonic/gin"
	"log"
)

type Test struct {
}

func NewTest() Test {
	return Test{}
}

func (t Test) Get(ctx *gin.Context) {
	log.Println("Hello 接收到GET请求..")
	response := app.NewResponse(ctx)
	response.ToResponse("接收GET请求成功")
}
