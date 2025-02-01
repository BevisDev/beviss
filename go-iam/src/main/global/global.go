package global

import (
	"goiam/src/main/config"

	"github.com/BevisDev/backend-template/database"
	"github.com/BevisDev/backend-template/logger"
	"github.com/BevisDev/backend-template/redis"
	"github.com/BevisDev/backend-template/rest"
)

var (
	AppConfig  *config.Config
	Logger     *logger.AppLogger
	AuthDB     *database.Database
	Redis      *redis.RedisClient
	RestClient *rest.RestClient
)
