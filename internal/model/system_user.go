package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type SystemUser struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Token    string `json:"token"`
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
	db = db.Where("username = ? AND is_del = ?", s.Username, 0)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func (s SystemUser) GetUserByToken(db *gorm.DB) (SystemUser, error) {
	var user SystemUser
	db = db.Where("token = ? AND is_del = ?", s.Token, 0)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func (s SystemUser) Update(db *gorm.DB, values interface{}) error {
	fmt.Println("更新内容", s.ID, values)
	if err := db.Model(&s).Where("id = ? AND is_del = ?", s.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t SystemUser) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Username != "" {
		db = db.Where("username = ?", t.Username)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t SystemUser) List(db *gorm.DB, pageOffset, pageSize int, usePage uint8) ([]*SystemUser, error) {
	var SystemUsers []*SystemUser
	var err error

	fmt.Println("usePage:", usePage)
	if usePage == 0 {
		if pageOffset >= 0 && pageSize > 0 {
			db = db.Offset(pageOffset).Limit(pageSize)
		}
	}

	if t.Username != "" {
		db = db.Where("username = ?", t.Username)
	}

	if err = db.Where("is_del = ?", 0).Find(&SystemUsers).Error; err != nil {
		return nil, err
	}
	return SystemUsers, nil
}
