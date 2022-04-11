package dao

import (
	"blog-service/internal/model"
	"time"
)

func (d *Dao) CreateSign(userId uint32, signTime string, createdBy string) error {
	signTimeFormat, _ := time.ParseInLocation("2006-01-06", signTime, time.Local)
	signIn := model.SignIn{
		UserId:   userId,
		SignTime: signTimeFormat,
		Model:    &model.Model{CreatedBy: createdBy},
	}
	_, error := signIn.Create(d.engine)
	return error
}
