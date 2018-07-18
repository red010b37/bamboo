package main

import (
	"log"
	"github.com/red010b37/bamboo/cmd/app/api"
	"github.com/red010b37/bamboo/cmd/app/conf"
	"github.com/red010b37/bamboo/cmd/app"
	"fmt"
	"runtime"
	"os"
)

// Idflags set by GoReleaser
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	// init app errors
	api.BuildAppErrors()

	// init rpc details
	conf.CreateRPCDetails()

	printInfo()


	// load the server config - required - contains server data
	err := conf.LoadServerConfig()
	if err != nil {
		log.Fatal("Failed to load the server config: " + err.Error())
	}

	// load the app config - required - contains active coin data
	err = conf.LoadAppConfig()
	if err != nil {
		log.Println("Failed to load the app config: " + err.Error())
	}

	// load the dev config file if one is set
	conf.LoadDevConfig()

	app.StartServer()
}

func printInfo()  {


	log.Println("--------------------------------")
	log.Println("  BAMBOO")
	log.Println(fmt.Sprintf("  Version: %v", version))
	log.Println(fmt.Sprintf("  Commit: %v", commit))
	log.Println(fmt.Sprintf("  Built: %v", date))
	log.Println("--------------------------------")

	// log out server runtime OS and Architecture
	log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
	log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))


}
