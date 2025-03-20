package startup

import (
	"github.com/BevisDev/backend-template/config"
	"github.com/BevisDev/backend-template/database"
	"github.com/BevisDev/backend-template/helper"
	"github.com/BevisDev/backend-template/logger"
	"github.com/BevisDev/backend-template/rest"
	"github.com/gin-gonic/gin"
	"gobaucua/src/main/lib"
	"gobaucua/src/main/router"
	"log"
	"sync"
)

func startup() *gin.Engine {
	var (
		state = helper.GenUUID()
		err   error
	)

	// get config
	err = config.GetConfig(&lib.AppConfig, "goauth", "yaml")
	if err != nil {
		log.Fatal("goauth config file read error: ", err)
		return nil
	}

	// startup logger
	lib.Logger = logger.NewLogger(&logger.ConfigLogger{
		Profile:    lib.AppConfig.ServerConfig.Profile,
		DirName:    lib.AppConfig.LoggerConfig.DirName,
		MaxSize:    lib.AppConfig.LoggerConfig.MaxSize,
		MaxBackups: lib.AppConfig.LoggerConfig.MaxBackups,
		MaxAge:     lib.AppConfig.LoggerConfig.MaxAge,
		Compress:   lib.AppConfig.LoggerConfig.Compress,
		IsSplit:    lib.AppConfig.LoggerConfig.IsSplit,
	})

	var wg sync.WaitGroup
	// rest client
	wg.Add(1)
	go func() {
		defer wg.Done()
		lib.RestClient = rest.NewRestClient(lib.AppConfig.ServerConfig.ClientTimeout, lib.Logger)
	}()

	// db
	wg.Add(1)
	go func() {
		defer wg.Done()
		// init auth db
		if lib.AuthDB, err = database.NewDB(&database.ConfigDB{
			Profile:      lib.AppConfig.ServerConfig.Profile,
			Kind:         helper.Postgres,
			Schema:       lib.AppConfig.Database.AuthDB.Schema,
			Host:         lib.AppConfig.Database.AuthDB.Host,
			Port:         lib.AppConfig.Database.AuthDB.Port,
			Username:     lib.AppConfig.Database.AuthDB.Username,
			Password:     lib.AppConfig.Database.AuthDB.Password,
			MaxOpenConns: lib.AppConfig.Database.AuthDB.MaxOpenConns,
			MaxIdleConns: lib.AppConfig.Database.AuthDB.MaxIdleConns,
			MaxLifeTime:  lib.AppConfig.Database.AuthDB.ConnMaxLifeTime,
			TimeoutSec:   lib.AppConfig.Database.AuthDB.TimeoutSec,
		}); err != nil {
			lib.Logger.Fatal(state, "authDB init error: {}", err)
		}
	}()

	// redis
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	if global.Redis, err = redis.NewRedis(&redis.RedisConfig{
	//		Host:     global.AppConfig.RedisConfig.Host,
	//		Port:     global.AppConfig.RedisConfig.Port,
	//		Password: global.AppConfig.RedisConfig.Password,
	//		DB:       global.AppConfig.RedisConfig.Index,
	//		PoolSize: global.AppConfig.RedisConfig.PoolSize,
	//	}); err != nil {
	//		global.Logger.Fatal(state, "redis init error: {}", err)
	//	}
	//}()

	wg.Wait()
	return router.InitRouter()
}
