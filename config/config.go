package config

import "time"

// StartupConfigDNS is a list of environment variables that needed to run DSN.
type StartupConfigDNS struct {
	SectorID        int64  `envconfig:"SECTOR_ID" required:"true"`
	DatabaseDSN     string `envconfig:"DATABASE_DSN" required:"true"`
	DatabaseMaxConn int32  `envconfig:"DATABASE_MAX_CONN" required:"true"`
	ListenAddr      string `envconfig:"LISTEN_ADDR" required:"true"`
}

// StartupConfigMigrator is a list of environment variables that needed to run Migrator.
type StartupConfigMigrator struct {
	DatabaseDSN      string        `envconfig:"DATABASE_DSN" required:"true"`
	MigrationPath    string        `envconfig:"MIGRATION_PATH" required:"true"`
	ConnectRetryTime time.Duration `envconfig:"CONNECT_RETRY_TIME" required:"true"`
}
