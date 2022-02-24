package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type CreateMessageRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	Content   string `form:"content" binding:"required,min=0,max=1000"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type MessageListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CountMessageRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//创建信息
func (svc *Service) CreateMessage(param *CreateMessageRequest) (uint32, error) {
	return svc.dao.CreateMessage(param.Title, param.Content, param.State, param.CreatedBy)
}

//接收信息
func (svc *Service) ReceiveMessage(param *CreateMessageRequest) error {
	return svc.dao.ReceiveMessage(param.Title, param.Content, param.State, param.CreatedBy)
}

func (svc *Service) CountMessage(param *CountMessageRequest) (int, error) {
	return svc.dao.CountMessage(param.Title, param.State)
}

func (svc *Service) GetMessageList(param *MessageListRequest, pager *app.Pager) ([]*model.Message, error) {
	return svc.dao.GetMessageList(param.Title, param.State, pager.Page, pager.PageSize)
}
