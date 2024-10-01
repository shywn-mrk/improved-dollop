package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"context"
)

func NewServer(lc fx.Lifecycle) *gin.Engine {
	server := gin.New()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.Run(":9000")
			return nil
		},
	})

	return server
}
