package dao

import "blog-service/internal/model"

func (d *Dao) CreateMessageTag(title string, state uint8, createdBy string) error {
	messageTag := model.MessageTag{
		Title: title,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	_, error := messageTag.Create(d.engine)
	return error
}
