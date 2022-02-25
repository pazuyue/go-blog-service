package dao

import (
	"blog-service/global"
	"blog-service/pkg/email"
)

func snedEmail(title string, count string) error {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return defailtMailer.SendMail(
		global.EmailSetting.To,
		title,
		count,
	)
}
