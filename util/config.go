package util

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DBName             string        `mapstructure:"DB_NAME"`
	ServerAddress      string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey  string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccesTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func Connect(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
