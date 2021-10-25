package main

import (
	"DNS/config"
	"DNS/logger"
	"database/sql"
	"expvar"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func main() {
	log := logger.Initialize()
	log.Info("Atlas Database Migrator Service start")

	cfg := config.StartupConfigMigrator{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	expvar.Publish("config_startup", expvar.Func(func() interface{} {
		return cfg
	}))

	db, err := sql.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		log.Info(err)
		os.Exit(1)
	}
	defer db.Close()

	for db.Ping() != nil {
		log.Info("waiting for db initialization...")
		time.Sleep(cfg.ConnectRetryTime)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Info("waiting for db initialization...")
		os.Exit(1)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(
		cfg.MigrationPath,
		"dsn", driver)
	if err != nil {
		log.Info(err)
		os.Exit(1)
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		log.Info(err)
		os.Exit(1)
	}
	log.Info("Migration process done!")
}
