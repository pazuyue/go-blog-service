package v1

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

//创建信息
func (t User) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateUserUser(&param)

	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToSuccessResponse(errcode.Success)
	return
}

//用户登录
func (t User) LoginByUserAndPassword(c *gin.Context) {
	param := service.LoginByUserAndPassword{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	param2 := service.AuthRequest{param.AppKey, param.AppSecret}
	err := svc.CheckAuth(&param2)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("svc.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	param.Token = token
	result := svc.LoginByUserAndPassword(&param)

	if !result {
		response.ToErrorResponse(errcode.ErrorLoginFail)
		return
	}

	response.ToResponse(gin.H{"code": errcode.Success.Code(), "data": gin.H{"message": errcode.Success.Msg(), "token": token}})
	return
}

//用户信息查询
func (t User) Info(c *gin.Context) {
	fmt.Println("Info")
	param := service.Info{}
	response := app.NewResponse(c)
	fmt.Println(response)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	result := svc.Info(&param)
	response.ToResponse(gin.H{"code": errcode.Success.Code(), "data": gin.H{"message": errcode.Success.Msg(), "roles": "admin", "user_id": result.ID, "user_name": result.Username, "avatar": "", "introduction": ""}})
}

//获取用户信息
func (t User) List(c *gin.Context) {
	param := service.UserList{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	SignList, err := svc.UserList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(SignList, totalRows)
	return
}
