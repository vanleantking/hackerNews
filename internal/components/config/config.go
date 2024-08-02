package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"POSTGRES_HOST"`
	DBPort                 string `mapstructure:"POSTGRES_PORT" default:"100"`
	DBUser                 string `mapstructure:"POSTGRES_USER"`
	DBPass                 string `mapstructure:"POSTGRES_PASS"`
	DBName                 string `mapstructure:"POSTGRES_DATABASE"`
	DBPoolIdle             int    `mapstructure:"POOL_IDLE" default:"10"`
	DBMaxConnection        int    `mapstructure:"MAX_CONNECTION" default:"100"`
	ConnLifeTime           int    `mapstructure:"CONN_LIFETIME" default:"300"`
	SSLMode                string `mapstructure:"SSL_MODE" default:"disable"`
	TimeZone               string `mapstructure:"TIMEZONE" default:"Asia/Ho_Chi_Minh"`
	LogLevel               int    `mapstructure:"LOG_LEVEL" default:"6"`
	ServerHost             string `mapstructure:"SERVER_HOST" default:"localhost"`
	APIPort                string `mapstructure:"API_PORT" default:"8080"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

var (
	once           sync.Once
	configInstance *Config
)

func NewConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't find the file .env : ", err)
		}

		err = viper.Unmarshal(&configInstance)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}

		if configInstance.AppEnv == "development" {
			log.Println("The App is running in development env")
		}
	})
	return configInstance
}
