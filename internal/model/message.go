package model

type Message struct {
	*Model
	TagID   uint32 `json:"tag_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a Message) TableName() string {
	return "blog_message"
}
