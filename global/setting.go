package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettinS
	CasbinSetting   *setting.CasbinSettingS
)
