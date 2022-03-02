package model

import (
	"fmt"
	"github.com/oddminng/go-blog-service/global"
	"github.com/oddminng/go-blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	logMode := logger.Default.LogMode(logger.Silent)
	if global.ServerSetting.RunMode == "debug" {
		logMode = logger.Default.LogMode(logger.Info)
	}
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{
		Logger: logMode,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqldb.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
