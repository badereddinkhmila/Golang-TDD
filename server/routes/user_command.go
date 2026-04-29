package routes

import (
	"go_tdd/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserCommandRoutes(router *gin.Engine) {
	{
		users := router.Group("/users")
		users.POST("/", handler.UserCreateCommand)
		users.DELETE("/:id", handler.UserDeleteCommand)
	}
}
