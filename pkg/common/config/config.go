package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

// Function that has a Predefined Return Value
func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev") // find file that named "dev" on ./pkg/common/config/envs
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
			return
	}

	err = viper.Unmarshal(&c)

	return
}