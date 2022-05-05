package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type CreateSignRequest struct {
	UserId    uint32 `json:"user_id" binding:"required,gte=1"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	SignTime  uint32 ` json:"sign_time" binding:"required"`
}

type SignListRequest struct {
	UserId uint32 `json:"user_id" `
	State  uint8  `form:"state" `
}

func (svc *Service) CreateSign(param *CreateSignRequest) error {
	return svc.dao.CreateSign(param.UserId, param.SignTime, param.CreatedBy)
}

func (svc *Service) CountSign(param *SignListRequest) (int, error) {
	return svc.dao.CountSign(param.UserId, param.State)
}

func (svc *Service) GetSignList(param *SignListRequest, pager *app.Pager) ([]*model.SignIn, error) {
	return svc.dao.GetSignList(param.UserId, param.State, pager.Page, pager.PageSize)
}
