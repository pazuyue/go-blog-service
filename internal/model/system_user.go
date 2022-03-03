package model

import "github.com/jinzhu/gorm"

type SystemUser struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s SystemUser) TableName() string {
	return "blog_system_user"
}

func (s SystemUser) Create(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s SystemUser) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", s.Model.ID, 0).Delete(&s).Error
}

func (s SystemUser) GetUser(db *gorm.DB) (SystemUser, error) {
	var user SystemUser
	db = db.Where("username = ? AND is_del = ?", s.Username, 1)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}
