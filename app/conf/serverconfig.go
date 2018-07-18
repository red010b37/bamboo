package conf

import (
	"github.com/red010b37/bamboo/app/utils"
	"github.com/spf13/viper"
)

// ServerConfig defines a structure to store server config data
type ServerConfig struct {
	ManagerAPIPort   int64
	DaemonAPIPort    int64
	SetupAPIPort     int64
	LatestReleaseAPI string
	ReleaseAPI       string
	DaemonHeartbeat  int64

	LivePort   int64
	TestPort   int64
	UseTestnet bool

	JWTSecret string // this is self generated on each start
}

// LoadServerConfig sets up viper, reads and parses server config
func LoadServerConfig() (ServerConfig, error) {

	viper.SetConfigName("server-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()

	if err != nil {
		return ServerConfig{}, err
	}

	serverConfig := parseServerConfig(ServerConfig{})

	ServerConf = serverConfig

	return serverConfig, nil

}

// parseServerConfig takes ServerConfig, parses and returns serverconf
func parseServerConfig(serverConf ServerConfig) ServerConfig {

	serverConf.ManagerAPIPort = viper.GetInt64("managerApiPort")
	serverConf.LatestReleaseAPI = viper.GetString("latestReleaseAPI")
	serverConf.ReleaseAPI = viper.GetString("releaseAPI")
	serverConf.DaemonHeartbeat = viper.GetInt64("daemonHeartbeat")

	serverConf.LivePort = viper.GetInt64("navCoinPorts.livePort")
	serverConf.TestPort = viper.GetInt64("navCoinPorts.testnetPort")
	serverConf.UseTestnet = viper.GetBool("useTestnet")

	return serverConf

}

func GenerateJWTSecret() {
	ServerConf.JWTSecret, _ = utils.GenerateRandomString(32)
}
