package dao

import (
	"blog-service/internal/model"
	"blog-service/internal/util"
)

func (d *Dao) CreateUserUser(username string, password string, createdBy string) error {
	systemUser := model.SystemUser{
		Username: username,
		Password: password,
		Model:    &model.Model{CreatedBy: createdBy},
	}

	return systemUser.Create(d.engine)
}

func (d *Dao) LoginByUserAndPassword(username string, password []byte) bool {
	systemUser := model.SystemUser{
		Username: username,
	}

	user, err := systemUser.GetUser(d.engine)
	if err != nil {
		return false
	}
	return util.ValidatePasswords(user.Password, password)
}

func (d *Dao) Info(username string) model.SystemUser {
	systemUser := model.SystemUser{
		Username: username,
	}
	user, _ := systemUser.GetUser(d.engine)
	return user

}
