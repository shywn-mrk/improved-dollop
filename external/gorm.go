package external

import (
	"context"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


func NewGorm(lc fx.Lifecycle) *gorm.DB {
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db := new(gorm.DB)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			db2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				return err
			}

			*db = *db2
			idb, err := db.DB()
			if err != nil {
				return err
			}

			return idb.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			if d, err := db.DB(); err == nil {
				return d.Close()
			} else {
				return err
			}
		},
	})

	return db
}
