package handler

import (
	"go_tdd/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserGetAllQuery(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	userService := service.NewUserSqlxService()
	users, err := userService.GetAllUser(ctx, uint(limit), uint(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"users": users})
}
