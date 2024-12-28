package util

import (
	"os"
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
	_, ci := os.LookupEnv("APP_CI")
	if !ci {
		viper.SetConfigFile(path + "/.env")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}
		err = viper.Unmarshal(&config)
	} else {
		config.DBName = os.Getenv("DB_NAME")
		config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
		config.AccesTokenDuration, _ = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	}
	return
}

func Connect(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
