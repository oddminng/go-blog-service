package global

import (
	"github.com/jinzhu/gorm"
	"github.com/oddminng/go-blog-service/pkg/logger"
	"github.com/oddminng/go-blog-service/pkg/setting"
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
