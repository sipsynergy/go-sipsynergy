package config

import (
	"bytes"
	"github.com/hashicorp/consul/api"
	"github.com/sipsynergy/go-sipsynergy/utils"
	"github.com/spf13/viper"
)

// LoadConfiguration determines the app config store (consul) from the environment and loads full configuration.
//
// Example env:
//   CONSUL_TOKEN = super_secret_token
//   CONSUL_HOST = localhost:8500
//   CONSUL_PATH = /services/feedback
//
func LoadConfiguration() *viper.Viper {
	v := viper.New()

	v.AutomaticEnv()
	config := api.DefaultConfig()
	config.Address = v.GetString("CONSUL_HOST")
	config.Token = v.GetString("CONSUL_TOKEN")

	c, err := api.NewClient(config)
	utils.PanicOnError(err, "failed to configure consul client")

	pair, _, err := c.KV().Get(v.GetString("CONSUL_PATH"), nil)
	utils.PanicOnError(err, "invalid consul config")

	if pair == nil || string(pair.Value) == "" {
		panic("received null configuration from consul")
	}

	v.SetConfigType("JSON")
	v.ReadConfig(bytes.NewBuffer(pair.Value))

	return v
}

