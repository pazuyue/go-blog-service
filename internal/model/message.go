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

func (t Message) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Title != "" {
		db = db.Where("title = ?", t.Title)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Message) List(db *gorm.DB, pageOffset, pageSize int) ([]*Message, error) {
	var messages []*Message
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db = db.Where("title = ?", t.Title)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
