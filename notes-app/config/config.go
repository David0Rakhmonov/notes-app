package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDSN string
	ServerPort  string
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetDefault("server.port", "8082")
	viper.SetDefault("database.dsn", "root:root@tcp(127.0.0.1:3306)/notes_app")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения файла конфигурации: %v", err)
	}

	AppConfig = &Config{
		ServerPort:  viper.GetString("server.port"),
		DatabaseDSN: viper.GetString("database.dsn"),
	}
	log.Println("Конфигурация успешно загружена")
}
