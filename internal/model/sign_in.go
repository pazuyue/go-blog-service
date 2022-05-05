package model

import (
	"github.com/jinzhu/gorm"
)

type SignIn struct {
	*Model
	UserId   uint32 `json:"userId"`
	SignTime uint32 `json:"sign_time"`
	State    uint8  `json:"state"`
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
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.UserId > 0 {
		db = db.Where("user_id = ?", t.UserId)
	}

	if t.State > 0 {
		db = db.Where("state = ?", t.State)
	}

	if err = db.Where("is_del = ?", 0).Find(&signIns).Error; err != nil {
		return nil, err
	}

	return signIns, nil
}
