package manager

import (
	"log"

	"github.com/red010b37/bamboo/app/conf"
	"github.com/red010b37/bamboo/app/daemon"
	"github.com/red010b37/bamboo/app/daemon/daemonapi"
	"github.com/gorilla/mux"
)

// StartAllDaemonManagers ranges through coins, starts daemons
func StartAllDaemonManagers(activeCoins []conf.CoinData) {

	log.Println("ranging through active coins, starting daemons")

	for _, coinData := range activeCoins {
		daemon.StartManager(coinData)
	}

}

// StartWalletHandlers ranges through activeCoins, inits handlers
func StartWalletHandlers(r *mux.Router, activeCoins []conf.CoinData) {

	log.Println("ranging through active coins, initialising wallet handlers")

	for _, coinData := range activeCoins {
		daemonapi.InitWalletHandlers(r, coinData, "api")
	}

}
