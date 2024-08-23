package components

import (
	"hackerNewsApi/pkg/config"
	"hackerNewsApi/pkg/gorm"
	"hackerNewsApi/pkg/logger"
	"hackerNewsApi/pkg/server"
	"hackerNewsApi/pkg/validator"
)

type App struct {
	DB        gorm.Database
	Logger    logger.Logger
	Validator *validator.Validator
	Server    server.Server
	Config    *config.Config
}

// for loading all components used in app
func AppConfig() *App {
	cfg := config.NewConfig()
	lg := logger.NewLogger(cfg)
	db := gorm.NewPostgresDatabase(cfg, lg)
	validate := validator.NewValidator(cfg)
	srv := server.NewServer(cfg)
	return &App{
		DB:        db,
		Logger:    lg,
		Validator: validate,
		Server:    srv,
		Config:    cfg,
	}
}
