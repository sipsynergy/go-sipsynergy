package utils

import (
	"bytes"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// LoadConfiguration determines the app config store (consul) from the environment and loads full configuration.
//
// Example env:
//   CONSUL_TOKEN = 0a47d149-a608-7b83-29c3-5b1f4d89e986
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
	PanicOnError(err, "")

	pair, _, err := c.KV().Get(v.GetString("CONSUL_PATH"), nil)
	PanicOnError(err, "Invalid consul config.")

	v.SetConfigType("JSON")
	v.ReadConfig(bytes.NewBuffer(pair.Value))

	return v
}

// PanicOnError logs error message and terminates the main process.
func PanicOnError(err error, message string) {
	if err != nil {
		HandleError(err, message)
	}
}
