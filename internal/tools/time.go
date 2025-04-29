package tools

import (
	"log"
	"time"
)

func GetDuration(lifetime string) time.Duration {
	d, err := time.ParseDuration(lifetime)
	if err != nil {
		log.Fatalln("Ошибка парсинга lifetime токена")
	}
	return d
}

func CheckExpire(t time.Time) bool {
	return time.Now().After(t)
}
