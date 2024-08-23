package migrator

import (
	"context"
	"hackerNewsApi/pkg/config"
	"hackerNewsApi/pkg/gorm"
	"hackerNewsApi/pkg/logger"

	"github.com/pressly/goose/v3"
)

const dialect = "postgres"

const (
	MigrateSeeder    = "seeder"
	MigrateMogration = "migrate"
)

type migrator struct {
	DB     gorm.Database
	logger logger.Logger
	cfg    config.Config
	cmd    string
	path   string
	args   []string
}

var ()

func NewMigrator(
	database gorm.Database,
	log logger.Logger,
	cfg config.Config,
	cmd string,
	path string,
	args []string,
) *migrator {
	return &migrator{
		DB:     database,
		logger: log,
		cfg:    cfg,
		cmd:    cmd,
		path:   path,
		args:   args}
}

func (m *migrator) Run() error {
	dbInstance, err := m.DB.GetDb().DB()
	if err != nil {
		return err
	}
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.RunContext(context.Background(), m.cmd, dbInstance, m.path, m.args...); err != nil {
		return err
	}
	return nil
}

func (m *migrator) Stop() error {
	return nil
}
