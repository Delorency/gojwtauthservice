package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigHTTPServer struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port string `env:"PORT" env-default:"8080"`
}

type ConfigDatabase struct {
	Role string `env:"DB_ROLE"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
}

type ConfigLogger struct {
	APIlp   string `env:"APILOGFILENAME"`
	DBlp    string `env:"DBLOGFILENAME"`
	LogsDir string `env:"LOGSDIR"`
}

type ConfigJWTToken struct {
	Typ string `env:"TYP"`
	Alg string `env:"ALG"`
	Iss string `env:"ISS"`

	SecretKey string `env:"SECRET_KEY"`

	Atl time.Duration
	Rtl time.Duration
}

type Config struct {
	HTTPServer *ConfigHTTPServer
	Db         *ConfigDatabase
	Logger     *ConfigLogger
	JWT        *ConfigJWTToken
}

func MustLoad() *Config {
	godotenv.Load("./configs/.env")

	var cfgHttpServer ConfigHTTPServer
	var cfgDatabase ConfigDatabase
	var cgfLogger ConfigLogger
	var cfgJWT ConfigJWTToken

	if err := cleanenv.ReadEnv(&cfgHttpServer); err != nil {
		log.Fatalln("Ошибка чтения настроек сервера из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgDatabase); err != nil {
		log.Fatalln("Ошибка чтения настроек бд из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cgfLogger); err != nil {
		log.Fatalln("Ошибка чтения настроек логгера из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgJWT); err != nil {
		log.Fatalln("Ошибка чтения настроек jwt токена из .env файлы")
	}

	accessDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalln("Ошибка парсинга lifetime access токена")
	}

	refreshDuration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalln("Ошибка парсинга lifetime access токена")
	}

	cfgJWT.Atl = accessDuration
	cfgJWT.Rtl = refreshDuration

	return &Config{&cfgHttpServer, &cfgDatabase, &cgfLogger, &cfgJWT}
}
