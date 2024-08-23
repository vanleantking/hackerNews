package cron

import (
	"hackerNewsApi/pkg/config"
	"hackerNewsApi/pkg/gorm"
	"hackerNewsApi/pkg/logger"
)

type App struct {
	DB     gorm.Database
	Logger logger.Logger
	Config *config.Config
}

// for loading all components used in app
func AppConfig() *App {
	cfg := config.NewConfig()
	lg := logger.NewLogger(cfg)
	db := gorm.NewPostgresDatabase(cfg, lg)
	return &App{
		DB:     db,
		Logger: lg,
		Config: cfg,
	}
}
