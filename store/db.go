package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"efbot/models"
)

var DB *gorm.DB

func Open() error {
	cfg := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	conn, err := gorm.Open(sqlite.Open(AppConfig.DBLocation()), cfg)
	if err != nil {
		return err
	}

	DB = conn

	return conn.AutoMigrate(&models.Message{}, &models.User{}, &models.Warning{})
}
