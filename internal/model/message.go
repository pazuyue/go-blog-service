package model

import "github.com/jinzhu/gorm"

type Message struct {
	*Model
	TagID   uint32 `json:"tag_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	State   uint8  `json:"state"`
}

func (a Message) TableName() string {
	return "blog_message"
}

func (t Message) Create(db *gorm.DB) (uint32, error) {
	result := db.Create(&t)
	return t.ID, result.Error
}

func (t Message) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Message) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
