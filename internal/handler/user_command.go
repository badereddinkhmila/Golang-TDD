package handler

import (
	"go_tdd/internal/domain"
	"go_tdd/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserCreateCommand(c *gin.Context) {
	ctx := c.Request.Context()

	var user *domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.NewUserSqlxService()
	user, err := userService.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func UserDeleteCommand(c *gin.Context) {
	ctx := c.Request.Context()
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userService := service.NewUserSqlxService()
	if err := userService.DeleteUser(ctx, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
