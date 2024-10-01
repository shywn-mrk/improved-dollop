package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lacion/mygolangproject/handler"
	"github.com/lacion/mygolangproject/service"
)

func SetupRoutes(engine *gin.Engine, service *service.Service) {
	engine.GET("/:id", handler.GetUser(service))
}
