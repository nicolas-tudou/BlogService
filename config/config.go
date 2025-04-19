package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Env       string `mapstructure:"ENV"`
	Port      string `mapstructure:"APP_PORT"`
	JWTSecret string `mapstructure:"JWTSECRET"`
	DB        struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
		SSLMode  string `mapstructure:"DB_SSL_MODE" default:"disable"`
		MaxConns int    `mapstructure:"DB_MAX_CONNS" default:"10"`
	} `mapstructure:"DB"`
}

var (
	AppConfig *Config
	once      sync.Once
)

func LoadConfig() {
	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error in read config file: %v\n", err)
		}

		if err := viper.Unmarshal(&AppConfig); err != nil {
			log.Fatal("Parse config file content error: %v\n", err)
		}
	})
}
