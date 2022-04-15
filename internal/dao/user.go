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

func (d *Dao) LoginByUserAndPassword(username string, password []byte, token string) bool {
	systemUser := model.SystemUser{
		Username: username,
	}

	user, err := systemUser.GetUser(d.engine)
	if err != nil {
		return false
	}

	values := map[string]interface{}{
		"token": token,
	}
	_ = user.Update(d.engine, values)
	return util.ValidatePasswords(user.Password, password)
}

func (d *Dao) Info(token string) model.SystemUser {
	systemUser := model.SystemUser{
		Token: token,
	}
	user, _ := systemUser.GetUserByToken(d.engine)
	return user

}
