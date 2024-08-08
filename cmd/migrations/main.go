package main

import (
	"flag"
	"fmt"
	"hackerNewsApi/internal/components/config"
	"hackerNewsApi/internal/components/gorm"
	"hackerNewsApi/internal/components/logger"
	"hackerNewsApi/internal/components/migrator"
)

var (
	strFlag string
)

func main() {
	flag.StringVar(&strFlag, "cmd", "", "help message")
	flag.Parse()

	fmt.Println("Non-flag arguments:", flag.Args())
	fmt.Println("-----")
	fmt.Println("strFlag value is: ", strFlag)
	cfg := config.NewConfig()

	lg := logger.NewLogger(cfg)
	db := gorm.NewPostgresDatabase(cfg, lg)

	migrate := migrator.NewMigrator(db, lg, *cfg, strFlag, flag.Args())
	fmt.Println(migrate.Run())
}
