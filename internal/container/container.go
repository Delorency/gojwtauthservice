package container

import (
	authdb "auth/internal/DB/authDB"
	userdb "auth/internal/DB/userDB"
	"auth/internal/config"
	service "auth/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	// Репозитории
	AuthRepo authdb.AuthDBI
	UserRepo userdb.UserDBI

	// Сервисы
	AuthService service.AuthServiceI
}

func NewContainer(db *gorm.DB, cfg *config.ConfigJWTToken, smtp *config.ConfigSMTP) *Container {
	userrepo := userdb.NewUserDB(db)
	authrepo := authdb.NewAuthDB(db, userrepo)
	authservice := service.NewAuthService(authrepo, userrepo, cfg, smtp)

	return &Container{
		AuthRepo:    authrepo,
		AuthService: authservice,
	}
}
