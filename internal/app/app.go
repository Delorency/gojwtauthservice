package app

import (
	"auth/internal/config"
	"auth/internal/container"
	"auth/internal/logger"
	"auth/internal/models"
	ht "auth/internal/transport/http"
	"auth/storage"
	"auth/storage/migrations"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Pass,
		DB:       cfg.Redis.DB,
	})

	container := container.NewContainer(db, rdb, cfg.JWT, cfg.WBhook)

	server := ht.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port, container, apilogger)

	CreateListTestUser(db) // создание тестовых пользователей

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		apilogger.Printf("Set IP: %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
		apilogger.Println("Starting server")

		log.Printf("Server work on %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("Server is stopping...")

	shtctx, shtcancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shtcancel()

	if err := server.Shutdown(shtctx); err != nil {
		log.Printf("Graceful shutdown error: %v\n", err)

		if err := server.Close(); err != nil {
			log.Fatalf("Forced termination error: %v\n", err)
		}
	}

}

func checkUpDB(logger gormlogger.Interface) *gorm.DB {
	db := storage.Psql(cfg.Db.Role, cfg.Db.Pass, cfg.Db.Name, cfg.Db.Host, cfg.Db.Port, logger)
	migrations.RunMigration(db)

	return db
}

func CreateListTestUser(db *gorm.DB) {
	u1 := &models.User{}
	u2 := &models.User{}
	u3 := &models.User{}
	db.Create(u1)
	db.Create(u2)
	db.Create(u3)
}
