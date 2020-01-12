package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/lucasstettner/api-boilerplate/db"
)

type (
	Config struct {
		Constants constants
		DB        *sql.DB
		Logger    *zap.Logger
	}

	constants struct {
		Postgres db.PostgresConfig `json:"postgres"`
	}
)

func New() *Config {
	config := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(dir + "/config.yml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	viper.Unmarshal(&config.Constants)
	if err != nil {
		panic(fmt.Errorf("Fatal error marshal config file: %s \n", err))
	}

	config.DB, err = db.CreateDatabase(&config.Constants.Postgres)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}

	config.Logger, err = zap.NewProduction()
	if err != nil {
		log.Panicf("Logging err: %v\n", err.Error())
	}

	return config
}
