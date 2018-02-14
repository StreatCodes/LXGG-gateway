package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/BurntSushi/toml"
)

//Config contains server configuration information
type Config struct {
	Addr      string
	LXDSocket string
	LXDAddr   string
}

//Load lxgg.toml
func loadConfig() Config {
	var config Config
	_, err := toml.DecodeFile("lxgg.toml", &config)

	if err != nil {
		log.Fatal("Could not open configuration file:", err)
		return Config{}
	}

	return config
}

//If the gateway is configured to use a unix socket we must return a custom http client
func createClient(config Config) http.Client {
	if config.LXDAddr == "http://unix" {
		return http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", config.LXDSocket)
				},
			},
		}
	}
	return http.Client{}
}
