package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ApplicationRouter() *gin.Engine {
	gin.ForceConsoleColor()

	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "gin.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
	}

	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
	UserCommandRoutes(router)
	UserQueryRoutes(router)

	return router
}
