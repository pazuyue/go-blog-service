package model

import (
	"github.com/jinzhu/gorm"
)

type SignIn struct {
	*Model
	SignTime uint32 `json:"sign_time"`
	State    uint8  `json:"state"`
	//外键
	UserId uint32 `json:"user_id"`
	User   User   `json:"systemUser";gorm:"foreignkey:UserId"` //指定关联外键
}

type User struct {
	ID       uint32 `gorm:"primary_key" json:"-"`
	Username string `json:"username"`
}

func (s User) TableName() string {
	return "blog_system_user"
}

func (s SignIn) TableName() string {
	return "blog_sign_in_info"
}

func (t SignIn) Create(db *gorm.DB) (uint32, error) {
	result := db.Create(&t)
	return t.ID, result.Error
}

func (t SignIn) Count(db *gorm.DB) (int, error) {
	var count int
	if t.UserId > 0 {
		db = db.Where("user_id = ?", t.UserId)
	}

	if t.State > 0 {
		db = db.Where("state = ?", t.State)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t SignIn) List(db *gorm.DB, pageOffset, pageSize int) ([]*SignIn, error) {
	var signIns []*SignIn
	var err error
	//var systemUser SystemUser
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.UserId > 0 {
		db = db.Where("user_id = ?", t.UserId)
	}

	if t.State > 0 {
		db = db.Where("state = ?", t.State)
	}
	db.Where("is_del = ?", 0)

	if err = db.Preload("User").Find(&signIns).Error; err != nil {
		return nil, err
	}

	return signIns, nil
}

func (t SignIn) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
