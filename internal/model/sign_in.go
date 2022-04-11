package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SignIn struct {
	*Model
	UserId   uint32    `json:"userId"`
	SignTime time.Time `json:"signTime"`
}

func (s SignIn) TableName() string {
	return "blog_sign_in_info"
}

func (t SignIn) Create(db *gorm.DB) (uint32, error) {
	result := db.Create(&t)
	return t.ID, result.Error
}
