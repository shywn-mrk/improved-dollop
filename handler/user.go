package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lacion/mygolangproject/service"
)

func GetUser(service *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")
		user, err := service.GetUser(userId)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	}
}
