package setting

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	DBHost   string `envconfig:"DBHost"`
	Username string `envconfig:"Username"`
	Password string `envconfig:"Password"`
	Database string `envconfig:"Database"`
}

type ServerConfig struct {
	HTTPPort     int           `envconfig:"PORT"`
	ReadTimeout  time.Duration `envconfig:"READTIMEOUT"`
	WriteTimeout time.Duration `envconfig:"WRITETIMEOUT"`
}

type AppConfig struct {
	RunMode   string `envconfig:"RUNMODE"`
	PageSize  int    `envconfig:"PAGESIZE"`
	JWTSecret string `envconfig:"JWTSECRET`
}

type Config struct {
	DBConfig
	ServerConfig
	AppConfig
}

func NewConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
