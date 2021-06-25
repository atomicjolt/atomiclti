package config

import (
	"encoding/base64"
	"encoding/json"
	"github.com/atomicjolt/atomiclti/lib"
	"log"
	"sync"
)

type ServerConfig struct {
	ServerPort       string `json:"server_port"`
	DbUser           string `json:"db_user"`
	DbHost           string `json:"db_host"`
	DbPassword       string `json:"db_password"`
	SessionSslMode   string `json:"session_ssl_mode"`
	SessionSecret    []byte `json:"session_secret"`
	Database         string `json:"database"`
	ClientId         string `json:"client_id"`
	AuthClientSecret []byte `json:"auth0_client_secret"`
	AuthClientId     string `json:"auth0_client_id"`
}

var once sync.Once
var cachedServerConfig *ServerConfig

func GetServerConfig() *ServerConfig {
	once.Do(func() {
		var configs map[string]ServerConfig
		err := json.Unmarshal(lib.LoadJsonFrom("./server_config.json"), &configs)
		if err != nil {
			log.Fatal("Server config file is not valid json: " + err.Error())
		}

		env := DetermineEnv()
		selectedConfig, isPresent := configs[env]
		if !isPresent {
			log.Fatal("Server config not found for env: " + env)
		}
		cachedServerConfig = &selectedConfig

		base64.StdEncoding.Decode(
			cachedServerConfig.SessionSecret,
			cachedServerConfig.SessionSecret,
		)

		base64.StdEncoding.Decode(
			cachedServerConfig.AuthClientSecret,
			cachedServerConfig.AuthClientSecret,
		)
	})

	return cachedServerConfig
}
