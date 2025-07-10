package app

import (
	"auth/internal/config"
	"auth/internal/container"
	"auth/internal/logger"
	"auth/internal/models"
	"auth/internal/transport/http"
	"auth/storage"
	"auth/storage/migrations"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var cfg *config.Config

func init() {
	cfg = config.MustLoad()
}

func Start() {
	apilogger := logger.GetAPILogger(
		fmt.Sprintf("%s/%s", cfg.Logger.LogsDir, cfg.Logger.APIlp),
	)
	dblogger := logger.GetDBLogger(
		fmt.Sprintf("%s/%s", cfg.Logger.LogsDir, cfg.Logger.DBlp),
	)

	db := checkUpDB(dblogger)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Pass,
		DB:       cfg.Redis.Name,
	})

	container := container.NewContainer(db, rdb, cfg.JWT, cfg.WBhook)

	server := http.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port, container, apilogger)

	apilogger.Printf("Set IP: %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	apilogger.Println("Starting server")

	log.Printf("Server work on %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	CreateListTestUser(db) // создание тестовых пользователей

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func checkUpDB(logger gormlogger.Interface) *gorm.DB {
	db := storage.Psql(cfg.Db.Role, cfg.Db.Pass, cfg.Db.Name, cfg.Db.Host, cfg.Db.Port, logger)
	migrations.RunMigration(db)

	return db
}

func CreateListTestUser(db *gorm.DB) {
	u1 := &models.User{TokenVersion: uuid.NewString(), Email: "user1@gmail.com"}
	u2 := &models.User{TokenVersion: uuid.NewString(), Email: "user2@gmail.com"}
	u3 := &models.User{TokenVersion: uuid.NewString(), Email: "user3@gmail.com"}
	db.Create(u1)
	db.Create(u2)
	db.Create(u3)
}
