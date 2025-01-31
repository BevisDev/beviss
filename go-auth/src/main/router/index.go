package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"goauth/src/main/config"
	"goauth/src/main/consts"
	"goauth/src/main/di"
	"goauth/src/main/dto/response"
	"goauth/src/main/middleware"
)

func RegisterRouter(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		AuthRouter(v1)
	}
}

func InitRouter() *gin.Engine {
	var r *gin.Engine
	serverConfig := config.AppConfig.ServerConfig

	if serverConfig.Profile == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// use Routers
	// ping to health check system
	pingC := di.NewPingDI()
	r.GET("/ping", pingC.Ping)
	r.GET("/db", pingC.PingDB)
	r.GET("/redis", pingC.PingRedis)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// use middleware
	middleware.RegisterMiddleware(r)

	// define group api
	api := r.Group("/api")
	{
		RegisterRouter(api)
	}

	// handler no route
	r.NoRoute(func(c *gin.Context) {
		response.SetError(c, 404, consts.NotFound)
	})
	return r
}
