package container

import (
	authdb "auth/internal/DB/authDB"
	userdb "auth/internal/DB/userDB"
	"auth/internal/config"
	service "auth/internal/service"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Container struct {
	// Репозитории
	AuthRepo authdb.AuthDBI
	UserRepo userdb.UserDBI

	// Сервисы
	AuthService service.AuthServiceI
}

func NewContainer(db *gorm.DB, rcfg *config.ConfigRedis, cfg *config.ConfigJWTToken, smtp *config.ConfigSMTP) *Container {
	userrepo := userdb.NewUserDB(db)
	authrepo := authdb.NewAuthDB(db, userrepo)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rcfg.Host, rcfg.Port),
		Password: rcfg.Pass,
		DB:       rcfg.Name,
	})

	authservice := service.NewAuthService(authrepo, rdb, userrepo, cfg, smtp)

	return &Container{
		AuthRepo:    authrepo,
		AuthService: authservice,
	}
}
