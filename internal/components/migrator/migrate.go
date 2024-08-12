package migrator

import (
	"context"
	"fmt"
	"hackerNewsApi/internal/components/config"
	"hackerNewsApi/internal/components/gorm"
	"hackerNewsApi/internal/components/logger"

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
	args   []string
}

var ()

func NewMigrator(
	database gorm.Database,
	log logger.Logger,
	cfg config.Config,
	cmd string,
	args []string,
) *migrator {
	return &migrator{
		DB:     database,
		logger: log,
		cfg:    cfg,
		cmd:    cmd,
		args:   args}
}

func (m *migrator) Run() error {
	dbInstance, err := m.DB.GetDb().DB()
	if err != nil {
		return nil
	}
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	args := m.args
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Println(args)
		return nil
	}
	path := m.cfg.PathMigrate
	if args[0] == MigrateSeeder {
		path += "/seeders"
		args = args[1:]
	} else {
		path += "/migrations"
	}

	if err := goose.RunContext(context.Background(), m.cmd, dbInstance, path, m.args...); err != nil {
		return err
	}
	return nil
}

func (m *migrator) Stop() error {
	return nil
}
