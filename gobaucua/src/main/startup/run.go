package startup

import (
	"gobaucua/src/main/lib"
)

func Run() {
	r := startup()
	defer func() {
		lib.Logger.SyncAll()
		lib.AuthDB.Close()
		lib.Redis.Close()
	}()

	// set trusted domain
	if err := r.SetTrustedProxies(lib.AppConfig.ServerConfig.TrustedProxies); err != nil {
		//logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(lib.AppConfig.ServerConfig.Port); err != nil {
		//logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
