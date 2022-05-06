package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

func (d *Dao) CreateSign(userId uint32, signTime uint32, createdBy string) error {
	signIn := model.SignIn{
		UserId:   userId,
		SignTime: signTime,
		Model:    &model.Model{CreatedBy: createdBy},
		State:    1,
	}
	_, error := signIn.Create(d.engine)
	return error
}

func (d *Dao) CountSign(user_id uint32, state uint8) (int, error) {
	signIn := model.SignIn{UserId: user_id, State: state}
	return signIn.Count(d.engine)
}

func (d *Dao) GetSignList(user_id uint32, state uint8, page, pageSize int) ([]*model.SignIn, error) {
	signIn := model.SignIn{UserId: user_id, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return signIn.List(d.engine, pageOffset, pageSize)
}
