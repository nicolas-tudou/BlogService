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
	// DB        DBConfig `mapstructure:"DB"`
	DB_Host     string `mapstructure:"DB_HOST"`
	DB_Port     string `mapstructure:"DB_PORT"`
	DB_User     string `mapstructure:"DB_USER"`
	DB_Password string `mapstructure:"DB_PASSWORD"`
	DB_Name     string `mapstructure:"DB_NAME"`
	DB_SSLMode  string `mapstructure:"DB_SSL_MODE" default:"disable"`
	DB_MaxConns int    `mapstructure:"DB_MAX_CONNS" default:"10"`
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
