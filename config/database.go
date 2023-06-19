package config

type Database struct {
	DatabaseURL  string `envconfig:"DATABASE_URL" required:"true"`
	MaxOpenConns int    `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"10"`
	MaxIdleConns int    `envconfig:"DATABASE_MAX_IDDLE_CONNECTIONS" default:"10"`
	MaxIdleTime  string `envconfig:"DATABASE_MAX_IDDLE_TIME" default:"10"`
}
