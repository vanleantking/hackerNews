package gorm

import (
	"fmt"
	"sync"
	"time"

	"hackerNewsApi/internal/components/config"

	appLog "hackerNewsApi/internal/components/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDatabase struct {
	Db *gorm.DB
}

type Database interface {
	GetDb() *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config, log appLog.Logger) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			conf.DBHost,
			conf.DBUser,
			conf.DBPass,
			conf.DBName,
			conf.DBPort,
			conf.SSLMode,
			conf.TimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.New(log, logger.Config{
				SlowThreshold:             time.Second * 5,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				LogLevel:                  logger.Info,
			}),
		})
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
