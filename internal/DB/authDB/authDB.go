package authdb

import (
	userdb "auth/internal/DB/userDB"
	"auth/internal/schemes"

	"gorm.io/gorm"
)

type AuthDBI interface {
	Access(*schemes.AccessCreate) (string, error)
}

type authDB struct {
	db     *gorm.DB
	userdb userdb.UserDBI
}

func NewAuthDB(db *gorm.DB, userdb userdb.UserDBI) AuthDBI {
	return &authDB{db, userdb}
}
