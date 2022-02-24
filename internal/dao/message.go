package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

func (d *Dao) CreateMessage(title string, content string, state uint8, createdBy string) (uint32, error) {
	message := model.Message{
		Title:   title,
		Content: content,
		State:   state,
		Model:   &model.Model{CreatedBy: createdBy},
	}

	return message.Create(d.engine)
}

func (d *Dao) ReceiveMessage(title string, content string, state uint8, createdBy string) error {
	tx := d.engine.Begin()

	messageTag := model.MessageTag{
		Title: title,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	tagId, err := messageTag.CreateOnlyNotExiste(d.engine)
	if err != nil {
		tx.Rollback()
		return err
	}

	message := model.Message{
		TagID:   tagId,
		Title:   title,
		Content: content,
		State:   state,
		Model:   &model.Model{CreatedBy: createdBy},
	}
	_, err2 := message.Create(tx)
	if err2 != nil {
		tx.Rollback()
		return err2
	}
	tx.Commit()
	return err2
}

func (d *Dao) CountMessage(title string, state uint8) (int, error) {
	message := model.Message{Title: title, State: state}
	return message.Count(d.engine)
}

func (d *Dao) GetMessageList(title string, state uint8, page, pageSize int) ([]*model.Message, error) {
	message := model.Message{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return message.List(d.engine, pageOffset, pageSize)
}
