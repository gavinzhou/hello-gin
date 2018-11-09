package setting

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	DBType     string `envconfig:"DBType"`
	DBHost     string `envconfig:"DBHost"`
	DBPort     string `envconfig:"DBPort"`
	DBUser     string `envconfig:"DBUser"`
	DBPassword string `envconfig:"DBPassword"`
	DBName     string `envconfig:"DBName"`
}

type ServerConfig struct {
	HTTPPort     int           `envconfig:"PORT"`
	ReadTimeout  time.Duration `envconfig:"READTIMEOUT"`
	WriteTimeout time.Duration `envconfig:"WRITETIMEOUT"`
}

type AppConfig struct {
	RunMode   string `envconfig:"RUN_MODE"`
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
