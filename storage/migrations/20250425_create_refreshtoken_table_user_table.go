package migrations

import (
	"auth/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateRefreshTokenTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250425_create_refreshtoken_table_user_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.RefreshToken{}, &models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(models.RefreshToken{}.TableName())
		},
	}
}
