package external

import (
	"context"

	"github.com/lacion/mygolangproject/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func InvokeMigrations(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.AutoMigrate(&models.User{}, &models.Address{})
		},
	})
}
