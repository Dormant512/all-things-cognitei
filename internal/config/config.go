package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	DBUser       string `env:"DB_USER" envDefault:"admin"`
	DBPassword   string `env:"DB_PASSWORD" envDefault:"pass"`
	DBName       string `env:"DB_NAME" envDefault:""`
	DBHost       string `env:"DB_HOST" envDefault:"localhost"`
	DBPort       string `env:"DB_PORT" envDefault:"27017"`
	DBType       string `env:"DB_TYPE" envDefault:"mongodb"`
	MGCollection string `env:"MG_COLLECTION" envDefault:"myCol"`
	MGSaveFile   string `env:"MG_COLLECTION" envDefault:"backup"`
	AppPort      string `env:"APP_PORT" envDefault:"3001"`
	DmUserJson   string `env:"DM_USER_JSON" envDefault:""`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg Config) GetConnectionString() string {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		cfg.DBType, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return dsn
}
