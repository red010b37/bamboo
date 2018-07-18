package app

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/red010b37/bamboo/cmd/app/manager"
	"github.com/red010b37/bamboo/cmd/app/api"
	"github.com/red010b37/bamboo/cmd/app/conf"
)

func StartServer(){

	// start the daemon managers for active coins
	manager.StartAllDaemonManagers(conf.AppConf.Coins)

	// setup the router
	router := mux.NewRouter()

	// setup the api meta and coin meta handlers
	api.InitMetaHandlers(router, "api")

	// start the transaction handlers for active coins
	manager.StartWalletHandlers(router, conf.AppConf.Coins)

	// set the proper server port
	port := fmt.Sprintf(":%d", conf.ServerConf.ManagerAPIPort)

	// start http server and listen up
	http.ListenAndServe(port, router)


}
