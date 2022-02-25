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

func (t MessageTag) Create(db *gorm.DB) (uint32, error) {
	result := db.Create(&t)
	return t.ID, result.Error
}

func (t MessageTag) CreateOnlyNotExiste(db *gorm.DB) (uint32, error) {
	db.First(&t, "title=?", t.Title)
	if t.ID > 0 {
		return t.ID, nil
	}
	result := db.Create(&t)
	return t.ID, result.Error
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

func (t MessageTag) GetOne(db *gorm.DB) (*MessageTag, error) {
	result := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).First(&t)
	return &t, result.Error
}
