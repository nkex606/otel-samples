package config

import "github.com/spf13/viper"

func InitConf() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			panic("Config.yaml file not found.")
		default:
			panic("Failed to load config.yaml.")
		}
	}
	return nil
}

type httpConfig struct {
	Port string `mapstructure:"port"`
}

type worldSvcConfig struct {
	Host string `mapstructure:"host"`
}

type otelConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

func GetHttpConfig() httpConfig {
	return httpConfig{
		Port: viper.GetString("http.port"),
	}
}

func GetWorldServerConfig() worldSvcConfig {
	return worldSvcConfig{
		Host: viper.GetString("world.host"),
	}
}

func GetOTELConfig() otelConfig {
	return otelConfig{
		Endpoint: viper.GetString("otel_collector.endpoint"),
	}
}

func GetSvcName() string {
	return viper.GetString("svcName")
}
