package global

import (
	"github.com/oddminng/go-blog-service/pkg/logger"
	"github.com/oddminng/go-blog-service/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

var (
	DBEngine *gorm.DB
	Logger   *logger.Logger
)
