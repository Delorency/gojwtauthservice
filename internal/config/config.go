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

type ConfigRedis struct {
	Host string `env:"REDIS_HOST"`
	Port int    `env:"REDIS_PORT"`
	Pass string `env:"REDIS_PASS"`
	Name int    `env:"REDIS_NAME"`
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

type ConfigSMTP struct {
	SmtpFrom string `env:"EMAIL_FROM"`
	SmtpPass string `env:"EMAIL_PASS"`
	SmtpHost string `env:"SMTPHOST"`
	SmtpPort string `env:"SMTPPORT"`
}

type ConfigWebhook struct {
	WBURL string `env:"WEBHOOKURL"`
}

type Config struct {
	HTTPServer *ConfigHTTPServer
	Db         *ConfigDatabase
	Redis      *ConfigRedis
	Logger     *ConfigLogger
	JWT        *ConfigJWTToken
	WBhook     *ConfigWebhook
}

func MustLoad() *Config {
	godotenv.Load("./configs/.env")

	var cfgHttpServer ConfigHTTPServer
	var cfgDatabase ConfigDatabase
	var cfgRedis ConfigRedis
	var cgfLogger ConfigLogger
	var cfgJWT ConfigJWTToken
	var cfgWebhook ConfigWebhook

	if err := cleanenv.ReadEnv(&cfgHttpServer); err != nil {
		log.Fatalln("Ошибка чтения настроек сервера из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgDatabase); err != nil {
		log.Fatalln("Ошибка чтения настроек бд из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgRedis); err != nil {
		log.Fatalln("Ошибка чтения настроек redis из .env файла")
	}
	if err := cleanenv.ReadEnv(&cgfLogger); err != nil {
		log.Fatalln("Ошибка чтения настроек логгеров из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgJWT); err != nil {
		log.Fatalln("Ошибка чтения настроек jwt токена из .env файлы")
	}
	if err := cleanenv.ReadEnv(&cfgWebhook); err != nil {
		log.Fatalln("Ошибка чтения настроек webhook из .env файлы")
	}

	accessDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalln("Ошибка парсинга lifetime access токена")
	}

	refreshDuration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalln("Ошибка парсинга lifetime refresh токена")
	}

	cfgJWT.Atl = accessDuration
	cfgJWT.Rtl = refreshDuration

	return &Config{&cfgHttpServer, &cfgDatabase, &cfgRedis, &cgfLogger, &cfgJWT, &cfgWebhook}
}
