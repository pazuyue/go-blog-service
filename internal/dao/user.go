package dao

import (
	"blog-service/internal/model"
	"blog-service/internal/util"
	"fmt"
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
	fmt.Println(user.Password)
	return util.ValidatePasswords(user.Password, password)
}
