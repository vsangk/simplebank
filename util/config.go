package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver  string `mapstructure:"DB_DRIVER"`
	DBSource  string `mapstructure:"DB_SOURCE"`
	DBAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // override with ENV variables if they exist

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}