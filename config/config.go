package config

import "github.com/spf13/viper"

type Config struct {
	ClientChartlyricsApi   string        `mapstructure:"CLIENT_CHARTLYRICS_API"`
	ClientAppleApi         string        `mapstructure:"CLIENT_APPLE_API"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}