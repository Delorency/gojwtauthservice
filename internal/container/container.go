package container

import (
	authdb "auth/internal/DB/authDB"
	userdb "auth/internal/DB/userDB"
	"auth/internal/config"
	service "auth/internal/service"

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

func NewContainer(db *gorm.DB, rdb *redis.Client, cfg *config.ConfigJWTToken, webhook *config.ConfigWebhook) *Container {
	userrepo := userdb.NewUserDB(db)
	authrepo := authdb.NewAuthDB(db, userrepo)

	authservice := service.NewAuthService(authrepo, rdb, userrepo, cfg, webhook)

	return &Container{
		AuthRepo:    authrepo,
		AuthService: authservice,
	}
}
