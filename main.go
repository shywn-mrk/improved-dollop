package main

import (
	"github.com/lacion/mygolangproject/external"
	"github.com/lacion/mygolangproject/repository"
	"github.com/lacion/mygolangproject/server"
	"github.com/lacion/mygolangproject/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			repository.NewAddressRepository,
			repository.NewUserRepository,
			service.NewService,
			external.NewGorm,
			server.NewServer,
		),
		fx.Invoke(
			external.InvokeMigrations,
			external.InvokeSeed,
			server.SetupRoutes,
		),
	).
	Run()
}
