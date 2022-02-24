package model

import "github.com/jinzhu/gorm"

type MessageTag struct {
	*Model
	Title string `json:"title"`
	State uint8  `json:"state"`
}

func (a MessageTag) TableName() string {
	return "blog_message_tag"
}

func (t MessageTag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t MessageTag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t MessageTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
