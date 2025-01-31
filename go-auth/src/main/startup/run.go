package startup

import (
	"goauth/src/main/infrastructure/database"
	"goauth/src/main/infrastructure/logger"
	"goauth/src/main/infrastructure/redis"
	"goauth/src/main/utils"
	"sync"
)

func Run() {
	// load configuration
	cf := startConfig()
	state := utils.GenUUID()

	// logger
	startLogger(state)

	// start app
	r := startRouter()

	var wg sync.WaitGroup

	// rest client
	wg.Add(1)
	go func() {
		defer wg.Done()
		startRestClient(state)
	}()

	// db
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	startDB(state)
	//}()

	// redis
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	startRedis(state)
	//}()

	wg.Wait()
	defer logger.SyncAll()
	defer database.CloseAll()
	defer redis.Close()

	// set trusted domain
	if err := r.SetTrustedProxies(cf.ServerConfig.TrustedProxies); err != nil {
		logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(cf.ServerConfig.Port); err != nil {
		logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
