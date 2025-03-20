package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gobaucua/src/main/consts"
	"gobaucua/src/main/di"
	"gobaucua/src/main/dto/response"
	"gobaucua/src/main/lib"
	"gobaucua/src/main/middleware"
)

func RegisterPublicAPIs(r *gin.Engine) {
	pingC := di.NewSupportDI()
	r.GET("/ping", pingC.Ping)
	r.GET("/db", pingC.PingDB)
	r.GET("/redis", pingC.PingRedis)
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterInternalAPIs(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		AuthRouter(v1)
	}
}

func RegisterWebSocketAPIs(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		AuthRouter(v1)
	}
}

func InitRouter() *gin.Engine {
	var r *gin.Engine
	serverConfig := lib.AppConfig.ServerConfig

	if serverConfig.Profile == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// public apis
	RegisterPublicAPIs(r)

	// use middleware
	middleware.RegisterMiddleware(r)

	// define group api
	api := r.Group("/api")
	{
		RegisterInternalAPIs(api)
	}

	// handler no route
	r.NoRoute(func(c *gin.Context) {
		response.SetError(c, 404, consts.NotFound)
	})
	return r
}
