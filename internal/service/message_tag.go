package service

type CreateMessageTagRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateMessageTag(param *CreateMessageTagRequest) error {
	return svc.dao.CreateMessageTag(param.Title, param.State, param.CreatedBy)
}
