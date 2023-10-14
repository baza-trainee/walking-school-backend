package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v9"
)

const (
	DebugLogLvl = "DEBUG"
	InfoLogLvl  = "INFO"
	ErrorLogLvl = "ERROR"
)

var errNotAllowedLoggelLevel = errors.New("not allowed logger level")

type Server struct {
	AppAddress      string        `env:"APP_PORT" envDefault:"7000"`
	AppReadTimeout  time.Duration `env:"APP_READ_TIMEOUT" envDefault:"60s"`
	AppWriteTimeout time.Duration `env:"APP_WRITE_TIMEOUT" envDefault:"60s"`
	AppIdleTimeout  time.Duration `env:"APP_IDLE_TIMEOUT" envDefault:"60s"`
}

type MongoDB struct {
	Host     string `env:"DB_HOST,notEmpty"`
	Port     string `env:"DB_PORT" envDefault:"27017"`
	DbName   string `env:"DB_NAME,notEmpty"`
	User     string `env:"DB_USER,notEmpty"`
	Password string `env:"DB_PASSWORD,notEmpty"`
	// SSLmode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}

type Form struct {
	Email         string `env:"FORM_EMAIL,notEmpty"`
	GeneratedCode string `env:"FORM_GENERATED_CODE,notEmpty"`
}

type Config struct {
	Server Server
	DB     MongoDB
	Form   Form
	// Auth     AuthConfig
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

// type AuthConfig struct {
// 	Salt            string        `env:"APP_SALT,notEmpty"`
// 	SigningKey      string        `env:"SIGNING_KEY,notEmpty"`
// 	AccessTokenTTL  time.Duration `env:"ACCESS_TOKEN_TTL" envDefault:"15m"`
// 	RefreshTokenTTL time.Duration `env:"REFRESH_TOKEN_TTL" envDefault:"24h"`
// }

func InitConfig() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("error while parsing .env: %w", err)
	}

	if err := validate(cfg.LogLevel); err != nil {
		return Config{}, fmt.Errorf("wrong loglevel value: %w", err)
	}

	cfg.Server.AppAddress = ":" + cfg.Server.AppAddress

	return cfg, nil
}

func validate(logLevel string) error {
	if strings.ToUpper(logLevel) != DebugLogLvl &&
		strings.ToUpper(logLevel) != ErrorLogLvl &&
		strings.ToUpper(logLevel) != InfoLogLvl {
		return errNotAllowedLoggelLevel
	}

	return nil
}
