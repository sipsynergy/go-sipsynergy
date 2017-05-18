package utils

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// LoadConfiguration determines the app config store (consul or etcd) from the environment and loads full configuration.
//
// Example env:
//   APP_CONFIG_PROVIDER = consul
//   APP_CONFIG_HOST = localhost:8500
//   APP_CONFIG_PATH = ./my-application-config.json
//
func LoadConfiguration() {
	viper.AutomaticEnv()

	viper.AddRemoteProvider(
		viper.GetString("APP_CONFIG_PROVIDER"),
		viper.GetString("APP_CONFIG_HOST"),
		viper.GetString("APP_CONFIG_PATH"),
	)

	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()
	PanicOnError(err, "Failed to read remote configuration.")
}

// PanicOnError logs error message and terminates the main process.
func PanicOnError(err error, message string) {
	if err != nil {
		HandleError(err, message)
	}
}