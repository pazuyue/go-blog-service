package model

type MessageTag struct {
	*Model
	Title string `json:"title"`
	State uint8  `json:"state"`
}

func (a MessageTag) TableName() string {
	return "blog_message_tag"
}
