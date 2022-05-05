package service

import (
	"blog-service/internal/model"
	"blog-service/internal/util"
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
