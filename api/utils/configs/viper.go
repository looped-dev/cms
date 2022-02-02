package configs

import "github.com/spf13/viper"

func init() {
	// configure viper to use environment variables or cms.yaml file
	viper.AddConfigPath(".")
	viper.SetConfigName("cms")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
}

// GetConfig gets a configuration value via viper. Viper uses either cms.yaml or
// environment variables to retrieve the value
func GetConfig(key string) string {
	return viper.GetString(key)
}

// SetConfig sets a configuration value via viper
func SetConfig(key, value string) {
	viper.Set(key, value)
}
