package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
	"fmt"
)

const (
	NOTHING = iota
	MAIL
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

//处理消息
func (d *Dao) MessageHandleRequest(state uint8, page, pageSize int) error {
	message := model.Message{State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	messages, err := message.List(d.engine, pageOffset, pageSize)
	if err != nil {
		return err
	}

	for _, v := range messages {
		tagId := *&v.TagID
		fmt.Println()
		MessageTag := model.MessageTag{Model: &model.Model{ID: tagId}}
		mt, _ := MessageTag.GetOne(d.engine)
		switch mt.State {
		case MAIL: //邮件发送
			err := snedEmail(v.Title, v.Content)
			if err != nil {
				return err
			}
			v.Update(d.engine, map[string]interface{}{"state": 1})
		}
	}

	return nil
}
