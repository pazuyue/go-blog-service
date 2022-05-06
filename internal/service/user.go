package service

import (
	"blog-service/internal/model"
	"blog-service/internal/util"
	"blog-service/pkg/app"
)

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=3,max=100"`
	Password  string `form:"password" binding:"required,min=10,max=1000"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type LoginByUserAndPassword struct {
	Username string `form:"username" binding:"required,min=3,max=100"`
	Password string `form:"password" binding:"required,min=10,max=1000"`
	Token    string
	AuthRequest
}

type Info struct {
	Token string `form:"token" binding:"required,min=3"`
}

type UserList struct {
	Username string `form:"username"`
	UsePage  uint8  `form:"use_page" binding:"oneof=0 1"`
}

//创建用户
func (svc *Service) CreateUserUser(param *CreateUserRequest) error {
	var Password []byte = []byte(param.Password)
	hashedPassword := util.HashAndSalt(Password)
	return svc.dao.CreateUserUser(param.Username, hashedPassword, param.CreatedBy)
}

//创建用户
func (svc *Service) LoginByUserAndPassword(param *LoginByUserAndPassword) bool {
	var Password []byte = []byte(param.Password)
	return svc.dao.LoginByUserAndPassword(param.Username, Password, param.Token)
}

//获取用户信息
func (svc *Service) Info(param *Info) model.SystemUser {
	return svc.dao.Info(param.Token)
}

//获取分页总数
func (svc *Service) CountUser(param *UserList) (int, error) {
	return svc.dao.CountUser(param.Username)
}

//获取分页
func (svc *Service) UserList(param *UserList, pager *app.Pager) ([]*model.SystemUser, error) {
	return svc.dao.UserList(param.Username, param.UsePage, pager.Page, pager.PageSize)
}
