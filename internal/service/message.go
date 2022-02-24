package service

type CreateMessageRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	Content   string `form:"content" binding:"required,min=0,max=1000"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) CreateMessage(param *CreateMessageRequest) (uint32, error) {
	return svc.dao.CreateMessage(param.Title, param.Content, param.State, param.CreatedBy)
}

func (svc *Service) ReceiveMessage(param *CreateMessageRequest) error {
	return svc.dao.ReceiveMessage(param.Title, param.Content, param.State, param.CreatedBy)
}
