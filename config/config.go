package config

import "time"

type StartupConfigDNS struct {
	DatabaseDSN     string `envconfig:"DATABASE_DSN" required:"true"`
	DatabaseMaxConn int32  `envconfig:"DATABASE_MAX_CONN" required:"true"`
	ListenAddr      string `envconfig:"LISTEN_ADDR" required:"true"`
}

type StartupConfigMigrator struct {
	DatabaseDSN      string        `envconfig:"DATABASE_DSN" required:"true"`
	MigrationPath    string        `envconfig:"MIGRATION_PATH" required:"true"`
	ConnectRetryTime time.Duration `envconfig:"CONNECT_RETRY_TIME" required:"true"`
}
