package main

import (
	"flag"
	"fmt"
	"hackerNewsApi/pkg/config"
	"hackerNewsApi/pkg/gorm"
	"hackerNewsApi/pkg/logger"
	"hackerNewsApi/pkg/migrator"
)

var (
	strFlag       string
	migrationPath string
	migrateType   string
)

func main() {
	flag.StringVar(&strFlag, "cmd", "", "command migrate")
	flag.StringVar(&migrationPath, "path", "", "data path file generate")
	flag.StringVar(&migrateType, "migrate", "", "migrate type is seeder of migration")
	flag.Parse()

	fmt.Println("Non-flag arguments:", flag.Args())
	fmt.Println("-----")
	fmt.Println("strFlag value is: ", strFlag, migrationPath)
	cfg := config.NewConfig()

	lg := logger.NewLogger(cfg)
	db := gorm.NewPostgresDatabase(cfg, lg)

	migrate := migrator.NewMigrator(db, lg, *cfg, strFlag, migrationPath, flag.Args())
	fmt.Println(migrate.Run())
}
