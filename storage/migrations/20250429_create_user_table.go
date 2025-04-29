package migrations

import (
	"auth/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250429_create_user_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.RefreshToken{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(models.RefreshToken{}.TableName())
		},
	}
}
