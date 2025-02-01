package startup

import (
	"goiam/src/main/global"
)

func Run() {
	r := start()
	defer func() {
		global.Logger.SyncAll()
		global.AuthDB.Close()
		global.Redis.Close()
	}()

	// set trusted domain
	if err := r.SetTrustedProxies(global.AppConfig.ServerConfig.TrustedProxies); err != nil {
		//logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(global.AppConfig.ServerConfig.Port); err != nil {
		//logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
