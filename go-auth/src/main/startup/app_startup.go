package startup

import (
	"github.com/gin-gonic/gin"
	"goauth/src/main/config"
	"goauth/src/main/router"
)

func startConfig() *config.Config {
	return config.LoadConfig()
}

func startRouter() *gin.Engine {
	return router.InitRouter()
}
