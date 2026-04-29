package routes

import (
	"go_tdd/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserQueryRoutes(router *gin.Engine) {
	{
		users := router.Group("/users")
		users.GET("/", handler.UserGetAllQuery)
	}
}
