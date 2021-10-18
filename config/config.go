package config

type StartupConfig struct {
	DatabaseDSN     string `envconfig:"DATABASE_DSN" required:"true"`
	DatabaseMaxConn int32  `envconfig:"DATABASE_MAX_CONN" required:"true"`
	ListenAddr      string `envconfig:"LISTEN_ADDR" required:"true"`
}
