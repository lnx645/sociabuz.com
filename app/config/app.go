package config

import (
	"log"

	"github.com/joho/godotenv"
	"sobuz.id/application/app/utils"
)

type Config struct {
	ServerPort string
	ServerHost string
	ServerEnv  string
}

var AppConfig *Config

func InitAppConfig() {

	if er := godotenv.Load(".env"); er != nil {
		log.Fatalf(".env Tidak ditemukan silahkan cobalagi! >>[%s]", er.Error())
	}
	AppConfig = &Config{
		ServerPort: utils.GetEnv("SERVER_PORT"),
		ServerHost: utils.GetEnv("SERVER_HOST"),
		ServerEnv:  utils.GetEnv("SERVER_ENV"),
	}
}
