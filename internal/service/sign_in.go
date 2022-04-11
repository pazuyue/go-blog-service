package service

type CreateSignRequest struct {
	UserId    uint32 `json:"user_id" binding:"required,gte=1"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	SignTime  string ` json:"sign_time" binding:"required"`
}

func (svc *Service) CreateSign(param *CreateSignRequest) error {
	return svc.dao.CreateSign(param.UserId, param.SignTime, param.CreatedBy)
}
