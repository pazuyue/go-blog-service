package dao

import (
	"blog-service/internal/model"
)

func (d *Dao) CreateSign(userId uint32, signTime uint32, createdBy string) error {
	signIn := model.SignIn{
		UserId:   userId,
		SignTime: signTime,
		Model:    &model.Model{CreatedBy: createdBy},
	}
	_, error := signIn.Create(d.engine)
	return error
}
