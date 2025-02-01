package startup

import (
	"github.com/BevisDev/backend-template/config"
	"github.com/BevisDev/backend-template/logger"
	"github.com/BevisDev/backend-template/rest"
	"github.com/gin-gonic/gin"
	"goiam/src/main/global"
	"goiam/src/main/router"
	"log"
	"sync"
)

func start() *gin.Engine {
	var (
		//state     = utils.GenUUID()
		err error
	)
	err = config.GetConfig(&global.AppConfig, "goauth", "yaml")
	if err != nil {
		log.Fatal("goauth config file read error: ", err)
		return nil
	}

	// start logger
	global.Logger = logger.NewLogger(&logger.ConfigLogger{
		Profile:    global.AppConfig.ServerConfig.Profile,
		FolderName: global.AppConfig.LoggerConfig.FolderName,
		MaxSize:    global.AppConfig.LoggerConfig.MaxSize,
		MaxBackups: global.AppConfig.LoggerConfig.MaxBackups,
		MaxAge:     global.AppConfig.LoggerConfig.MaxAge,
		Compress:   global.AppConfig.LoggerConfig.Compress,
		IsSplit:    global.AppConfig.LoggerConfig.IsSplit,
	})

	var wg sync.WaitGroup

	// rest client
	wg.Add(1)
	go func() {
		defer wg.Done()
		global.RestClient = rest.NewRestClient(global.AppConfig.ServerConfig.ClientTimeout)
	}()

	// db
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	// init auth db
	//	if global.AuthDB, err = database.NewDB(&database.ConfigDB{
	//		Profile:      global.AppConfig.ServerConfig.Profile,
	//		Kind:         "postgres",
	//		Schema:       global.AppConfig.Database.AuthDB.Schema,
	//		Host:         global.AppConfig.Database.AuthDB.Host,
	//		Port:         global.AppConfig.Database.AuthDB.Port,
	//		Username:     global.AppConfig.Database.AuthDB.Username,
	//		Password:     global.AppConfig.Database.AuthDB.Password,
	//		MaxOpenConns: global.AppConfig.Database.AuthDB.MaxOpenConns,
	//		MaxIdleConns: global.AppConfig.Database.AuthDB.MaxIdleConns,
	//		MaxLifeTime:  global.AppConfig.Database.AuthDB.ConnMaxLifeTime,
	//		TimeoutSec:   global.AppConfig.Database.AuthDB.TimeoutSec,
	//	}); err != nil {
	//		global.Logger.Fatal(state, "authDB init error: {}", err)
	//	}
	//}()

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
